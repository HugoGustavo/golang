check_environment_variables() {
    local required_vars=("$1")

    for var in "${required_vars[@]}"; do
        if [ -z "${!var}" ]; then
            echo "Error: $var is not set."
            return 1
        fi
    done

    echo "Environment variables configured properly."
    return 0
}
