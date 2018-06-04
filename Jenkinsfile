pipeline {
    agent { docker {
        image 'golang:1.10.2-stretch'
        customWorkspace '/go/src/bibucket.org/andreacaban/hello'
    } }
    environment {
        APP_DIR = "$GOPATH/src/bibucket.org/andreacaban/hello"
    }
    stages {
        stage('setup') {
            steps {
                sh 'chown -R jenkins /go'
                sh 'go version'
            }
        }
        stage('build') {
            steps {
                sh 'echo $APP_DIR'
                sh 'echo $GOPATH'
                sh 'pwd && ls -al'
                sh 'go build -v && ls -al'
            }
        }
        stage('check') {
            steps {
                sh 'pwd'
                sh 'ls -al'
                sh 'echo $GOPATH'
                sh 'pwd'
            }
        }
    }
}