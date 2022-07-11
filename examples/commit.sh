#!/bin/bash

# This script is used to write a conventional commit message.
# It prompts the user to choose the type of commit as specified in the
# conventional commit spec. And then prompts for the summary and detailed
# description of the message and uses the values provided. as the summary and
# details of the message.
#
# If you want to add a simpler version of this script to your dotfiles, use:
#
# alias gcm='git commit -m "$(gum input)" -m "$(gum write)"'


TYPE=$(gum choose "fix" "feat" "docs" "style" "refactor" "test" "chore" "revert")
SCOPE=$(gum input --placeholder "scope")

if [ -n "$SCOPE" ]; then
    SCOPE="($SCOPE)"
fi

git commit \
    -m "$(gum input --value "$TYPE$SCOPE: " --placeholder "Summary of this change")" \
    -m "$(gum write --placeholder "Details of this change")"
