import groovy.json.JsonSlurper;

node {
    def committer
    def groupId = "nais"
    def appConfig = "nais.yaml"
    def application = "coregroups"
    def dockerDir = "./docker"
    def distDir = "${dockerDir}/dist"
    def go = "/usr/local/go/bin/go"

    stage("checkout") {
	    git url: "https://github.com/navikt/${application}.git"
    }

    try {
	    stage("initialize") {
        releaseVersion = sh(script: 'echo $(date "+%Y-%m-%d")-$(git --no-pager log -1 --pretty=%h)', returnStdout: true).trim()
        committer = sh(script: 'git log -1 --pretty=format:"%ae (%an)"', returnStdout: true).trim()
      }

      stage("compile binary and prepare build") {
        sh "CGO_ENABLED=0 GOOS=linux ${go} build -a -installsuffix cgo -o coregroups ."
      }

      stage("build and publish docker image") {
        def imageName = "docker.adeo.no:5000/${application}:${releaseVersion}"
        sh "sudo docker build -t ${imageName} ."
        sh "sudo docker push ${imageName}"
      }

      stage("publish yaml") {
        withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'nexusUser', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
		sh "curl -s -F r=m2internal -F hasPom=false -F e=yaml -F g=${groupId} -F a=${application} -F v=${releaseVersion} -F p=yaml -F file=@${appConfig} -u ${env.USERNAME}:${env.PASSWORD} http://maven.adeo.no/nexus/service/local/artifact/maven/content"
        }
      }

      stage("deploy to !prod") {
        withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'srvauraautodeploy', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
                sh "curl -k -d \'{\"application\": \"${application}\", \"version\": \"${releaseVersion}\", \"fasitEnvironment\": \"cd-u1\", \"zone\": \"fss\", \"namespace\": \"default\", \"fasitUsername\": \"${env.USERNAME}\", \"fasitPassword\": \"${env.PASSWORD}\"}\' https://daemon.nais.preprod.local/deploy"
        }
      }
      
      stage("verify resources") {
	      retry(15) {
		      sleep 5
          httpRequest consoleLogResponseBody: true,
                      ignoreSslErrors: true,
                      responseHandle: 'NONE',
                      url: 'https://coregroups.nais.preprod.local/isalive',
                      validResponseCodes: '200'
	      }
      }

      stage("deploy to prod") {
        withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'srvauraautodeploy', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
          sh "curl -k -d \'{\"application\": \"${application}\", \"version\": \"${releaseVersion}\", \"fasitEnvironment\": \"p\", \"zone\": \"fss\", \"namespace\": \"default\", \"fasitUsername\": \"${env.USERNAME}\", \"fasitPassword\": \"${env.PASSWORD}\"}\' https://daemon.nais.adeo.no/deploy"
        }
      }

      slackSend channel: '#nais-ci', message: ":nais: Successfully deployed ${application}:${releaseVersion} to prod :partyparrot: \nhttps://${application}.nais.adeo.no\nLast commit by ${committer}.", teamDomain: 'nav-it', tokenCredentialId: 'slack_fasit_frontend'

      if (currentBuild.result == null) {
        currentBuild.result = "SUCCESS"
      }
    } catch(e) {
      if (currentBuild.result == null) {
        currentBuild.result = "FAILURE"
      }

      slackSend channel: '#nais-ci', message: ":shit: Failed deploying ${application}:${releaseVersion}: ${e.getMessage()}. See log for more info ${env.BUILD_URL}", teamDomain: 'nav-it', tokenCredentialId: 'slack_fasit_frontend'
      throw e
    } finally {
        step([$class       : 'InfluxDbPublisher',
              customData   : null,
              customDataMap: null,
              customPrefix : null,
              target       : 'influxDB'])
    }
}
