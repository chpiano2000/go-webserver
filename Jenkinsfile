pipeline {
  environment {
      dockerHubCredential = '43fb44ba-27e9-4061-8db0-d20dbe3689c6'
      imageName = 'datvc/go-webserver'
      containerName = 'recipe'
      dockerImage = ''
      tag = ''
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
          def commitHash = sh('git log -1 --pretty=%h')
          tag = '${imageName}:${commitHash}'
          echo tag
          echo '$tag'
          dockerImage = docker.build "${imageName}:${commitHash}"
        }
      }
    }

    stage('Push Image') {
      steps {
        script {
          // Use Jenkins credentials for Docker Hub login
          withCredentials([usernamePassword(credentialsId: dockerHubCredential, usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
              sh "docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD"

              // Push the image
              sh "docker push $tag"
          }
        }
      }
    }

  }
}
