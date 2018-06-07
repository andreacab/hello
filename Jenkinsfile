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
            }
        }
        stage('build') {
            steps {
                sh 'cd $APP_DIR && go build -v'
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