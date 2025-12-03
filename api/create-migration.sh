#!/usr/bin/env bash

export PATH="$PATH:$HOME/go/bin"

show_help() {
    cat <<EOF
Usage: $0 <migration_name> [OPTION]

Options:
  -h, --help       Show this help message.

Description:
  A simple bash script for creating migrations using goose.
EOF
}

create_migration() {
    local migration_name="$1"
    if [[ -z "$migration_name" ]]; then
        echo "Error: migration name is required."
        exit 1
    fi

    goose -s -dir ./internal/db/migrations create "$migration_name" sql
}

main() {
    case "$1" in
        -h|--help)
            show_help
            ;;
        "")
            echo "Error: no migration name provided."
            echo "Use --help for usage information."
            exit 1
            ;;
        *)
            create_migration "$1"
            ;;
    esac
}

main "$@"
