import Vue from "vue";
import Router from "vue-router";
import Home from "./views/Home.vue";
import Login from "./views/Login.vue";
import ListServices from "./views/ServiceDiscovery/ListServices.vue";
import NewService from "./views/ServiceDiscovery/NewService.vue";
import ViewService from "./views/Service/ViewService.vue";
import ByApi from "./views/Analytics/ByApi.vue";

var OAuthValidator = require("@/auth");
Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: "home",
      component: Home
    },
    {
      path: "/login",
      name: "login",
      component: Login
    },
    {
      path: "/service-discovery/services",
      name: "service-discovery-services",
      component: ListServices
    },
    {
      path: "/service-discovery/services/new",
      name: "service-discovery-services-new",
      component: NewService,
      beforeEnter: OAuthValidator.requireAuth
    },
    {
      path: "/service-discovery/service",
      name: "service-view",
      component: ViewService
    },
    {
      path: "/analytics/by-api",
      name: "analytics-by-api",
      component: ByApi,
      beforeEnter: OAuthValidator.requireAuth
    }
  ]
});
