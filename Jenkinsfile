pipeline {
    agent {
        docker {
            image 'argo.registry:5000/epel-7-mgo1.15'
            args '-u jenkins:jenkins'
        }
    }
    options {
        checkoutToSubdirectory('ams-push-worker')
        newContainerPerStage()
    }
    environment {
        PROJECT_DIR="ams-push-worker"
        GH_USER = 'newgrnetci'
        GH_EMAIL = '<argo@grnet.gr>'
        GOPATH="${WORKSPACE}/go"
        GIT_TAG=sh(returnStdout: true, script: "git tag --sort version:refname | tail -1").trim()
        GIT_COMMIT=sh(script: "cd ${WORKSPACE}/$PROJECT_DIR && git log -1 --format=\"%H\"",returnStdout: true).trim()
        GIT_COMMIT_HASH=sh(script: "cd ${WORKSPACE}/$PROJECT_DIR && git log -1 --format=\"%H\" | cut -c1-7",returnStdout: true).trim()
        GIT_COMMIT_DATE=sh(script: "date -d \"\$(cd ${WORKSPACE}/$PROJECT_DIR && git show -s --format=%ci ${GIT_COMMIT_HASH})\" \"+%Y%m%d%H%M%S\"",returnStdout: true).trim()
    }
    stages {
        stage('Binary Builds') {
            steps {
                
                sh """
                mkdir -p ${WORKSPACE}/go/src/github.com/ARGOeu
                ln -sf ${WORKSPACE}/${PROJECT_DIR} ${WORKSPACE}/go/src/github.com/ARGOeu/${PROJECT_DIR}
                rm -rf ${WORKSPACE}/go/src/github.com/ARGOeu/${PROJECT_DIR}/${PROJECT_DIR}
                cd ${WORKSPACE}/go/src/github.com/ARGOeu/${PROJECT_DIR}
                export CGO_CFLAGS"=-O2 -fstack-protector --param=ssp-buffer-size=4 -D_FORTIFY_SOURCE=2"
                go build -buildmode=pie -ldflags "-s -w -linkmode=external -extldflags '-z relro -z now'"
                /home/jenkins/checksec.py -b ./ams-push-worker
                zip ams-push-worker_linux_x86_64.zip ./ams-push-worker
                """
                archiveArtifacts artifacts: 'ams-push-worker/*.zip'
            }
            post{
                always {
                    cleanWs()
                }
            }
        }
        
    }
    post{
        always {
            cleanWs()
        }
    }
}