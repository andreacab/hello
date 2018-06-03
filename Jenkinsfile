pipeline {
    agent { docker { image 'golang:1.10.2-stretch' } }
    stages {
        stage('setup') {
            steps {
                sh 'go version'
            }
        }
        stage('build') {
            steps {
                sh 'mkdir -p $GOPATH/src/bibucket.org/andreacaban/hello'
                sh 'mv -v ./* $GOPATH/src/bibucket.org/andreacaban/hello'
                sh 'cd $GOPATH/src/bibucket.org/andreacaban/hello && pwd && ls -al'
                sh 'pwd'
                sh 'cd $GOPATH/src/bibucket.org/andreacaban/hello && go build -v'
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