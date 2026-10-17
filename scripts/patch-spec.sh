#!/bin/bash
# Downconverts Roark's OpenAPI 3.1 schemas for oapi-codegen's 3.0 reader:
# nullable type arrays become nullable:true, const becomes a one-value enum,
# schema examples becomes example, and numeric exclusive bounds become 3.0
# boolean flags with minimum/maximum. Null union branches become nullable,
# and single-branch unions are collapsed to avoid undefined N0 accessors.
# A nullable reference uses allOf so the 3.0 reader retains nullable instead
# of ignoring it as a sibling of $ref.
# All paths and operations are retained.
set -euo pipefail
trap 'on_error "$?" "$LINENO"' ERR

LOG_FILE="${LOG_FILE:-/tmp/$(basename "$0" .sh).log}"
exec > >(tee -a "$LOG_FILE") 2>&1

SPEC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/api"
readonly SPEC_DIR
readonly SPEC_IN="${SPEC_DIR}/openapi.json"
readonly SPEC_OUT="${SPEC_DIR}/openapi.fixed.json"

on_error() {
	local exit_code="$1"
	local line="$2"

	printf '{"level":"ERROR","file":"%s","line":%d,"msg":"command failed","exit_code":%d}\n' \
		"${BASH_SOURCE[1]}" "$line" "$exit_code" >&2
}

log() {
	local level="$1"
	shift

	jq -nc \
		--arg time "$(date -u '+%Y-%m-%dT%H:%M:%SZ')" \
		--arg level "$level" \
		--arg file "${BASH_SOURCE[1]##*/}" \
		--arg line "${BASH_LINENO[0]}" \
		--arg func "${FUNCNAME[1]:-main}" \
		--arg msg "$*" \
		'{time:$time,level:$level,file:$file,line:($line|tonumber),func:$func,msg:$msg}' >&2
}

main() {
	local tmp
	tmp="$(mktemp "${SPEC_DIR}/.openapi.fixed.XXXXXX")"
	# shellcheck disable=SC2064 # expand tmp before the local goes out of scope
	trap "rm -f -- '${tmp}'" EXIT

	jq '
	  def fix_union($key):
	    if (.[$key]? | type) == "array" then
	      (if any(.[$key][]; .type? == "null") then
	        .[$key] |= map(select(.type? != "null")) | .nullable = true
	      else . end)
	      | if (.[$key] | length) == 1 then
	          if .nullable == true and (.[$key][0] | has("$ref")) then
	            .allOf = .[$key] | del(.[$key])
	          else .[$key][0] + del(.[$key]) end
	        else . end
	    else . end;
	  walk(if type == "object" then
	    (if (.type? | type) == "array" then
	      if (.type | length) == 2 and (.type | index("null")) != null then
	        .type = (.type - ["null"] | .[0]) | .nullable = true
	      else error("unsupported type array") end
	    else . end)
	    | fix_union("oneOf") | fix_union("anyOf")
	    | (if has("const") then .enum = [.const] | del(.const) else . end)
	    | (if (.examples? | type) == "array" then
	        .example = (.example // .examples[0]) | del(.examples)
	      else . end)
	    | (if (.exclusiveMinimum? | type) == "number" then
	        .minimum = .exclusiveMinimum | .exclusiveMinimum = true
	      else . end)
	    | (if (.exclusiveMaximum? | type) == "number" then
	        .maximum = .exclusiveMaximum | .exclusiveMaximum = true
	      else . end)
	  else . end)
	  | .openapi = "3.0.3"
	' "$SPEC_IN" >"$tmp"
	chmod 0644 "$tmp"
	mv -- "$tmp" "$SPEC_OUT"
	log INFO "downconverted spec out=${SPEC_OUT##*/}"
}

main "$@"
