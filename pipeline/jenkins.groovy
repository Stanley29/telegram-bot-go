pipeline {
    agent any

    parameters {
        choice(name: 'OS', choices: ['linux', 'windows', 'darwin'], description: 'Target operating system')
        choice(name: 'ARCH', choices: ['amd64', 'arm64'], description: 'Target architecture')
        booleanParam(name: 'SKIP_TESTS', defaultValue: false, description: 'Skip tests')
        booleanParam(name: 'SKIP_LINT', defaultValue: false, description: 'Skip linter')
    }

    stages {
        stage('Clone') {
            steps {
                echo "Cloning for ${params.OS}/${params.ARCH}"
            }
        }
        stage('Lint') {
            when {
                expression { return !params.SKIP_LINT }
            }
            steps {
                echo "Running linter..."
            }
        }
        stage('Test') {
            when {
                expression { return !params.SKIP_TESTS }
            }
            steps {
                echo "Running tests..."
            }
        }
        stage('Build') {
            steps {
                echo "Building for ${params.OS}/${params.ARCH}"
            }
        }
        stage('Push') {
            steps {
                echo "Pushing image..."
            }
        }
    }
}
