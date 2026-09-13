pipeline {
    agent any

    environment {
        APP_NAME = 'multibranch-go-app'
        REPO_URL = "https://github.com/RahmaMohamed205/multibranch-go-app.git"
    }

    stages {

        stage('Getting Repo files') {
            steps {
                git branch: "${GIT_BRANCH}", credentialsId: 'github-creds', url: "${REPO_URL}"
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
                          -e BRANCH_NAME=${GIT_BRANCH} \
                          -p 0:5000 \
                          ${APP_NAME}:${BUILD_NUMBER}
                    """
                }
            }
        }

        stage('Verify App Is Running') {
            steps {
                script {
                    def hostPort = sh(
                        script: "docker port ${APP_NAME}-test 5000/tcp | cut -d: -f2",
                        returnStdout: true
                    ).trim()

                    echo "App running at http://localhost:${hostPort}"
                    sh "sleep 2"
                    sh "curl -f http://localhost:${hostPort}/health"
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
}
