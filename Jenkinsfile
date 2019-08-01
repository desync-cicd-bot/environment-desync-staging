pipeline {
  options {
    disableConcurrentBuilds()
  }
  agent {
    label "jenkins-go"
  }
  environment {
    DEPLOY_NAMESPACE = "jx-staging"
  }
  stages {
    stage('Validate Environment') {
      steps {
        container('go') {
          dir('env') {
            sh 'jx step helm build'
          }
        }
      }
    }
    stage('Update Environment') {
      when {
        branch 'master'
      }
      steps {
        container('go') {
          dir('env') {
            sh 'jx step helm apply'
          }
        }
      }
    }
    stage('Test') {
        when {
            branch 'master'
        }
        steps {
            container('go') {
                sh 'make test'
            }
        }
    }
  }
}
