pipeline {
    agent any

    environment {
        APP_NAME = 'multibranch-go-app'
        REPO_URL = "https://github.com/RahmaMohamed205/multi-branch.git"
    }

    stages {

        stage('Getting Repo files') {
            steps {
                git branch: "${env.BRANCH_NAME}", credentialsId: 'github-creds', url: "${REPO_URL}"
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh """
                        docker build -t ${APP_NAME}:${BUILD_NUMBER} .
                    """
                }
            }
        }

        stage('Check Image Size') {
            steps {
                sh "docker images ${APP_NAME}:${BUILD_NUMBER}"
            }
        }

        stage('Push Docker Image') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                        sh """
                            echo \$DOCKER_PASSWORD | docker login -u \$DOCKER_USERNAME --password-stdin
                            docker tag ${APP_NAME}:${BUILD_NUMBER} \$DOCKER_USERNAME/${APP_NAME}:${BUILD_NUMBER}
                            docker push \$DOCKER_USERNAME/${APP_NAME}:${BUILD_NUMBER}
                        """
                    }
                }
            }
        }

        stage('Run Container') {
    steps {
        script {
            sh "docker rm -f ${APP_NAME}-test || true"
            sh """
                docker run -d \
                  --name ${APP_NAME}-test \
                  -e BRANCH_NAME=${env.BRANCH_NAME} \
                  ${APP_NAME}:${BUILD_NUMBER}
            """
        }
    }
}

        stage('Verify App Is Running') {
    steps {
        script {
            def containerIp = sh(
                script: "docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' ${APP_NAME}-test",
                returnStdout: true
            ).trim()

            echo "App running at http://${containerIp}:5000"

            sh "sleep 2"
            sh "curl -f http://${containerIp}:5000/health"
        }
    }
}

        stage('Cleanup') {
            steps {
                sh "docker rm -f ${APP_NAME}-test || true"
            }
        }
    }
}
