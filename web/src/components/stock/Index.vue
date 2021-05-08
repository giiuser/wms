<template>
    <section>
        <b-field label="Начните вводить наименование товара">
            <b-autocomplete
                v-model="name"
                :data="data"
                placeholder="например хлеб бородинский"
                field="title"
                :loading="isFetching"
                @input="getAsyncData"
                @select="option => selected = option">

                <template slot-scope="props">
                    <div class="media">
                        <div class="media-content">
                            {{ props.option.name }}
                        </div>
                    </div>
                </template>
            </b-autocomplete>
        </b-field>
        <div class="card-content">
            <div class="content">
                <p class>
                    Доступные остатки товара {{ productName }}:
                    <strong>{{ Number(totalStock) }}</strong>
                </p>
            </div>
            <hr />
            <h6>Наличие на полках:</h6>
            <div class="card-content">
                <div class="content" v-for="(cell, key) in cellstocks" :key="key">
                    <p class>
                        <strong>{{ cell.cell_name }}:</strong>
                        {{ cell.qty }}
                    </p>
                </div>
            </div>
        </div>
    </section>
</template>

<script>
import Axios from "axios";
import * as debounce from 'lodash/debounce'

Axios.defaults.baseURL = 'http://localhost:8010';

export default {
    name: "StockIndex",
        computed: {},
        data: function() {
            return {
                data: [],
                name: '',
                productName: '',
                selected: null,
                isFetching: false,
                totalStock: null,
                cellstocks: []
            };
        },
        watch: {
            selected: function (val) {
            Axios.get(`/stock/${val.id}`)
                    .then(({ data }) => {
                        this.productName = val.name
                        this.totalStock = data.qty
                        this.cellstocks = data.cells
                        console.log(data)
                    })
                    .catch((error) => {
                        console.log(error)
                    })
            }
        },
        methods: {
            getAsyncData: debounce(function () {
                this.data = []
                this.isFetching = true
                Axios.get(`/productsearch/?query=${this.name}`)
                    .then(({ data }) => {
                        data.forEach((item) => this.data.push(item))
                        console.log(this.data)
                        this.isFetching = false
                    })
                    .catch((error) => {
                        this.isFetching = false
                        throw error
                    })
            }, 500)
        },
    }
</script>