<template>
    <div class="wares-table">
        <section>
            <b-field>
                <b-button tag="router-link" to="/collect/new" type="is-success">Создать новый</b-button>
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
                    <b-table-column field="created_at" label="Создан" v-slot="props">{{formatDate(props.row.created_at)}}</b-table-column>
                    <b-table-column field="created_at" label="Изменен" v-slot="props">{{formatDate(props.row.updated_at)}}</b-table-column>
                    <b-table-column field="status" label="Статус" v-slot="props">
                        <b-tag :type="props.row.status ? 'is-success' : 'is-warning'" v-if="props.row.status !== 3">{{props.row.status ? 'проведен' : 'не проведен'}}</b-tag>
                        <b-tag type="is-dark" v-else>в архиве</b-tag>
                    </b-table-column>
                    <b-table-column field="actions" label="Действия" v-slot="props">
                        <div class="buttons">
                        <router-link :to="'/allocation/' + props.row.id" v-if="props.row.status !== 3" class="button is-info is-small">Редактирование</router-link>
                        <b-button v-if="props.row.status === 0" type="is-danger" @click="deleteRow(props.row.id)" size="is-small">Удалить</b-button>
                        <b-button v-else-if="props.row.status !== 3" type="is-warning" @click="setArchive(props.row.id)" size="is-small">В архив</b-button>
                        </div>
                    </b-table-column>
            </b-table>
        </section>
    </div>
</template>

<script>
import Axios from "axios";
import moment from 'moment';

// Axios.defaults.baseURL = 'http://localhost:8010';

export default {
    name: "CollectIndex",
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

            Axios.get("/collects")
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
        deleteRow: async function (id) {
            await Axios.delete('/collect/' + id);
            this.loadAsyncData();
        },
        formatDate: function(value) {
            console.log(value);
            if (value) {
                return moment(String(value)).format('DD-MM-YYYY hh:mm')
            }
        },
        setArchive: async function (id) {
            await Axios.patch('/collect/' + id,
                {
                    status: 3
                }
            )
            this.loadAsyncData();
        },
    },
    mounted() {
        this.loadAsyncData();
    }
};
</script>