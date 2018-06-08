pipeline {
    agent {
        docker {
            image 'golang:1.10.2-stretch'
            args  '-v $WORKSPACE:/go/src/bibucket.org/andreacaban/hello -e "APP_DIR=/go/src/bibucket.org/andreacaban/hello"'
            reuseNode true
        }
    }
    stages {
        stage('test app') {
            steps {
                sh 'echo "Testing Go app"'
                sh 'cd $APP_DIR && go test'
            }
        }
        stage('Build app') {
            steps {
                sh 'echo "Installing dependencies"'
                sh 'cd $APP_DIR && go get'
                sh 'echo "Building Go app"'
                sh 'cd $APP_DIR && go build -v'
            }
        }
        stage('build image') {
            agent none
            steps {
                steps {
                    sh "docker build -t andreacab/hello:${GIT_COMMIT} ."
                    sh "docker image ls"

                }
            }
        }
        stage('check image runs') {
            agent none
            steps {
                sh 'cd $APP_DIR && ls -al'
                sh 'echo "Done!"'
            }
        }
    }
}