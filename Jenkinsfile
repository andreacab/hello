pipeline {
    agent { docker {
        image 'golang:1.10.2-stretch'
        args  '-v $WORKSPACE:/go/src/bibucket.org/andreacaban/hello -e "APP_DIR=/go/src/bibucket.org/andreacaban/hello"'
        reuseNode true
    } }
    stages {
        stage('setup') {
            steps {
                sh 'go version && echo $APP_DIR'
                sh 'git log -n 1 --pretty=format:"%h"'
            }
        }
        stage('build app') {
            steps {
                sh 'cd $APP_DIR && go build -v'
            }
        }
        stage('build image') {
            steps {
                sh '''
                    docker image ls
                '''
            }
        }
        stage('check') {
            steps {
                sh 'cd $APP_DIR && ls -al'
                sh 'echo "Done!"'
            }
        }
    }
}