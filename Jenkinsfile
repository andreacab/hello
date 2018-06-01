pipeline {
    agent { docker { image 'golang:1.10.2-stretch' } }
    stages {
        stage('test setup') {
            steps {
                sh 'go version'
            }
        }
        stage('build') {
            steps {
                sh 'go build'
            }
        }
        stage('check') {
            steps {
                sh 'pwd'
                sh 'ls -al'
            }
        }
    }
}