import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';
import ProductIndex from './components/product/Index';
import ProductCreate from './components/product/Create';
import ProductEdit from './components/product/Edit';
import ReceiptIndex from './components/receipt/Index';
import ReceiptEdit from './components/receipt/Edit';
import AllocationIndex from './components/allocation/Index';
import AllocationEdit from './components/allocation/Edit';
import CollectIndex from './components/collect/Index';
import CollectEdit from './components/collect/Edit';
import CellIndex from './components/cell/Index';
import CellCreate from './components/cell/Create';
import WaybillIndex from './components/waybill/Index';
import WaybillEdit from './components/waybill/Edit';
import StockIndex from './components/stock/Index';

import Axios from "axios";

// Axios.defaults.baseURL = 'http://localhost:8010';

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
    { path: '/cells', component: CellIndex },
    {
         path: "/cells/new",
         name: "createcell",
         component: CellCreate,
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
        async beforeEnter(routeTo, routeFrom, next) {
            try {
                const response = await Axios.post("/receipt", {});

                if (response.status !== 201) {
                    throw "posting error";
                }

                next("/receipt/" + response.data.id);
            } catch (error) {
                alert(error);
            }
        }
    },
    {
        path: "/receipt/:receiptId",
        name: "editreceipt",
        component: ReceiptEdit
    },
    { path: '/allocations', component: AllocationIndex },
    {
        path: "/allocation/new",
        name: "createallocation",
        component: AllocationEdit,
        async beforeEnter(routeTo, routeFrom, next) {
            try {
                const response = await Axios.post("/allocation", {});

                if (response.status !== 201) {
                    throw "posting error";
                }

                next("/allocation/" + response.data.id);
            } catch (error) {
                alert(error);
            }
        }
    },
    {
        path: "/allocation/:allocationId",
        name: "editallocation",
        component: AllocationEdit
    },
    { path: '/collects', component: CollectIndex },
    {
        path: "/collect/new",
        name: "createcollect",
        component: CollectEdit,
        async beforeEnter(routeTo, routeFrom, next) {
            try {
                const response = await Axios.post("/collect", {});

                if (response.status !== 201) {
                    throw "posting error";
                }

                next("/collect/" + response.data.id);
            } catch (error) {
                alert(error);
            }
        }
    },
    {
        path: "/collect/:collectId",
        name: "editcollect",
        component: CollectEdit
    },
    { path: '/waybills', component: WaybillIndex },
    {
        path: "/waybill/new",
        name: "createwaybill",
        component: WaybillEdit,
        async beforeEnter(routeTo, routeFrom, next) {
            try {
                const response = await Axios.post("/waybill", {});

                if (response.status !== 201) {
                    throw "posting error";
                }

                next("/waybill/" + response.data.id);
            } catch (error) {
                alert(error);
            }
        }
    },
    {
        path: "/waybill/:waybillId",
        name: "editwaybill",
        component: WaybillEdit
    },
    { path: '/stocks', component: StockIndex },
 ]
});