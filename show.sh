set -eux pipefail

find .

RUNFILES="bazel-bin/show.runfiles/__main__"
if [ ! -d "${RUNFILES}" ]; then
    RUNFILES="."
fi

PRESENT_TOOL="${RUNFILES}/external/org_golang_x_tools/cmd/present/linux_amd64_stripped/present" 
#PRESENT_BASE="${RUNFILES}/external/org_golang_x_tools_cmd_present"
PRESENT_BASE="${RUNFILES}"

$PRESENT_TOOL -base $PRESENT_BASE -notes



