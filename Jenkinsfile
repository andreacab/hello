pipeline {
    agent { docker { image 'golang:1.10.2-stretch' } }
    stages {
        stage('build') {
            steps {
                sh 'echo "hello world"'
                sh 'go --version'
            }
        }
    }
}