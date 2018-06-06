pipeline {
    agent { docker {
        image 'andreacab/hello:latest'
        args  '-v $WORKSPACE:/go/src/bibucket.org/andreacaban/hello'
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