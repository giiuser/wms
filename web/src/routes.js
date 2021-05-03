import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';
import ProductIndex from './components/product/Index';
import ProductCreate from './components/product/Create';
import ProductEdit from './components/product/Edit';
import ReceiptIndex from './components/receipt/Index';
import ReceiptEdit from './components/receipt/Edit';

import Axios from "axios";

Axios.defaults.baseURL = 'http://localhost:8010';

Vue.use(VueRouter);

export default new VueRouter({
 mode: 'history',
 routes: [
   { path: '/', component: Home },
   { path: '/products', component: ProductIndex },
   {
        path: "/products/new",
        name: "createproduct",
        component: ProductCreate,
        // async beforeEnter(routeTo, routeFrom, next) {
        //     try {
        //         const response = await Axios.post("/product", {
        //             stock_id: 1
        //         });

        //         if (response.status !== 200) {
        //             throw "delete error";
        //         }

        //         console.log(response.data);

        //         next("/products/" + response.data);
        //     } catch (error) {
        //         alert(error);
        //     }
        // }
    },
    {
        path: "/products/:productId",
        name: "editproduct",
        component: ProductEdit,
        // beforeEnter(routeTo, routeFrom, next) {
        //     if (store.getters["auth/isAuthenticated"]) {
        //         store
        //             .dispatch("request/fetchRequest", {
        //                 id: routeTo.params.docId
        //             })
        //             .then(response => {
        //                 next();
        //             })
        //             .catch(() => {
        //                 //next({ name: '404', params: { resource: 'User' } })
        //             });
        //     }
        // }
    },
    { path: '/receipts', component: ReceiptIndex },
    {
        path: "/receipt/new",
        name: "createreceipt",
        component: ReceiptEdit,
        // async beforeEnter(routeTo, routeFrom, next) {
        //     try {
        //         const response = await Axios.post("/receipt");

        //         if (response.status !== 200) {
        //             throw "posting error";
        //         }

        //         next("/receipt/" + response.data);
        //     } catch (error) {
        //         alert(error);
        //     }
        // }
    },
    // {
    //     path: "/receipt/:receiptId",
    //     name: "editreceipt",
    //     component: ReceiptEdit
    // },
 ]
});