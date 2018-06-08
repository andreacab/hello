pipeline {
    agent {
        docker {
            image 'golang:1.10.2-stretch'
            args  '-v $WORKSPACE:/go/src/bibucket.org/andreacaban/hello -e "APP_DIR=/go/src/bibucket.org/andreacaban/hello"'
            reuseNode true
        }
    }
    stages {
        stage('setup app') {
            steps {
                sh 'cd $APP_DIR && go list ./...'
                sh 'cd $APP_DIR && go get ./...'
            }
        }
        stage('test app') {
            steps {
                sh 'echo "Testing Go app"'
                sh 'cd $APP_DIR && go test'
            }
        }
        stage('Build app') {
            steps {
                sh 'echo "Building Go app"'
                sh 'cd $APP_DIR && go build -v'
            }
        }
        stage('build image') {
            agent none
            steps {
                script {
                    docker.build("andreacab/hello:${env.GIT_COMMIT}")
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