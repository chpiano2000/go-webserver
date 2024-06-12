pipeline {
  environment {
      dockerHubCredential = '43fb44ba-27e9-4061-8db0-d20dbe3689c6'
      imageName = 'datvc/go-webserver'
      containerName = 'recipe'
      dockerImage = ''
  }

  agent any

  stages {

    stage('Clone Repository') {
      steps {
        echo 'Cloning repositories'
        git([url: 'https://github.com/chpiano2000/go-webserver.git', branch: 'cicd'])
      }
    }

    stage('Build image') {
      steps{
        echo 'Starting to build docker image'
        script {
          def commitHash = sh(script: 'git log -1 --pretty=%h', returnStdout: true)
          env.COMMIT_HASH = commitHash
        }
        echo '${imageName}:${COMMIT_HASH}'
        script {
          dockerImage = docker.build "${imageName}:${COMMIT_HASH}"
        }
      }
    }

    // stage('Push Image') {
    //  steps {
    //    script {
    //      // Use Jenkins credentials for Docker Hub login
    //      withCredentials([usernamePassword(credentialsId: dockerHubCredential, usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
    //          sh "docker push ${imagename}:${COMMIT_HASH}"
    //      }
    //    }
    //  }
    // }

  }
}
