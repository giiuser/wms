<template>
    <div class="wares-table">
        <section>
            <b-field>
                <b-button tag="router-link" to="/products/new" type="is-success">Создать новый</b-button>
            </b-field>
            <b-table
                :data="data"
                :loading="loading"
                paginated
                :total="total"
                :per-page="perPage"
                @page-change="onPageChange"
            >
                    <b-table-column field="id" label="ID" v-slot="props">{{props.row.id}}</b-table-column>
                    <b-table-column field="name" label="Наименование" v-slot="props">{{props.row.name}}</b-table-column>
                    <b-table-column field="brand" label="Производитель" v-slot="props">{{props.row.brand}}</b-table-column>
                    <b-table-column field="actions" label="Действия" v-slot="props">
                        <router-link :to="'/products/' + props.row.id" class="button is-info is-small">Просмотр</router-link>
                    </b-table-column>
            </b-table>
        </section>
    </div>
</template>

<script>
import Axios from "axios";

Axios.defaults.baseURL = 'http://localhost:8010';

export default {
    name: "Index",
    computed: {},
    data: function() {
        return {
            data: [],
            total: 0,
            loading: false,
            page: 1,
            perPage: 20,
            selected: {},
        };
    },
    methods: {
        loadAsyncData() {
            this.loading = true;

            Axios.get("/products")
                .then(response => {
                    let currentTotal = response.data.length;
                    if (response.data.length / this.perPage > 1000) {
                        currentTotal = this.perPage * 1000;
                    }
                    this.total = currentTotal;
                    this.loading = false;
                    this.data = response.data;
                })
                .catch(error => {
                    console.log(error);
                });
        },
        onPageChange(page) {
            this.page = page;
            this.loadAsyncData();
        },
        editTodo(obj) {
            console.log(obj);
            obj.editing = true;
        },
        doneEdit(obj) {
            obj.editing = false;
        },
        pullRow(row) {
            this.$emit("update-selected", row);
        }
    },
    mounted() {
        this.loadAsyncData();
    }
};
</script>