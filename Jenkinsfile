pipeline {

  agent any

  stages {

    stage('Checkout Source') {
      steps {
        git 'https://github.com/chpiano2000/go-webserver.git'
      }
    }

    stage('Test') {
      steps{
        script {
          sh echo 'Success'
        }
      }
    }


  }

}
