build:
	CGO_ENABLED=0 GOOS=linux go build .

containers: build
	docker build -t hotrod .

ocp: ocp-build ocp-deploy

ocp-deploy: ocp-deploy-frontend

ocp-deploy-frontend:
	 oc create -f deployments/openshift/frontend.yaml ; \
	 oc expose dc frontend --port 8080 ; \
	 oc expose service frontend

ocp-deploy-customer:
	 oc create -f deployments/openshift/customer.yaml ; \
	 oc expose dc customer --port 8080 ; \
	 oc expose service customer

ocp-deploy-driver:
	 oc create -f deployments/openshift/driver.yaml ; \
	 oc expose dc driver --port 8080 ; \
	 oc expose service driver

ocp-deploy-route:
	 oc create -f deployments/openshift/route.yaml ; \
	 oc expose dc route --port 8080 ; \
	 oc expose service route

ocp-build-route: build
	 oc new-build --name route --binary ; \
	 oc start-build route --from-dir=. --follow 	 

ocp-build-frontend: build
	 oc new-build --name frontend --binary ; \
	 oc start-build frontend --from-dir=. --follow 
	
ocp-build-customer: build
	 oc new-build --name customer --binary ; \
	 oc start-build customer --from-dir=. --follow 

ocp-build-driver: build
	 oc new-build --name driver --binary ; \
	 oc start-build driver --from-dir=. --follow 
