<template>
    <div class="wares-table">
        <section>
            <b-field grouped group-multiline>
                <b-input></b-input>
                <router-link to="/postings/new" class="button is-success">Создать оприходование</router-link>
            </b-field>
            <b-table
                :data="data"
                :loading="loading"
                paginated
                backend-pagination
                :total="total"
                :per-page="perPage"
                :selected.sync="selected"
                @page-change="onPageChange">
                <template slot-scope="props">
                    <b-table-column field="id" label="ID">{{props.row.id}}</b-table-column>
                    <b-table-column field="stock_id" label="Склад">{{props.row.stock.name}}</b-table-column>
                    <b-table-column field="created_at" label="Создан">{{props.row.created_at}}</b-table-column>
                    <b-table-column field="status" label="Статус">
                        <b-tag :type="props.row.status ? 'is-success' : 'is-danger'">{{props.row.status ? 'проведен' : 'не проведен'}}</b-tag>
                    </b-table-column>
                    <b-table-column field="actions" label="Действия">
                        <router-link :to="'/postings/' + props.row.id" class="button is-info is-small">Просмотр</router-link>
                        <button class="button is-success is-small" v-if="props.row.allocation">Размещено</button>
                        <button class="button is-warning is-small" @click="makeAllocation(props.row.id, props.row.stock_id)" v-else :disabled="!props.row.status">Разместить</button>
                    </b-table-column>
                </template>
            </b-table>
        </section>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex'
    import Axios from 'axios'

    export default {
        name: "Posting",
        data: function() {
            return {
                data: [],
                total: 0,
                loading: false,
                page: 1,
                perPage: 20,
                selected: {}
            }
        },
        methods:{
            loadAsyncData() {
                this.loading = true

                Axios.get('/api/postings').then(response => {
                    let currentTotal = response.data.length
                    if (response.data.length / this.perPage > 1000) {
                        currentTotal = this.perPage * 1000
                    }
                    this.total = currentTotal
                    this.loading = false
                    this.data = response.data
                }).catch(error => {
                    console.log(error)
                })
            },
            onPageChange(page) {
                this.page = page
                this.loadAsyncData()
            },
            makeAllocation(id, stock_id) {
                this.$store.dispatch('allocation/setDocument', {document_id: id, document_type: 'postings'})
                this.$store.dispatch('allocation/setStock', stock_id)
                this.$router.push('/allocations/new')
            },
        },
        mounted() {
            this.loadAsyncData()
        }
    }
</script>

<style scoped>

</style>