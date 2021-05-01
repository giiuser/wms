<template>
    <div class="content">
        <loading :active.sync="isLoading" :is-full-page="fullPage"></loading>
        <div class="card">
            <div class="card-content">
                <div class="content">
                    <div class="content">
                        <b-select placeholder="На какой склад" v-model="toStock">
                            <option v-for="stock in stocks" :key="stock.id" :value="stock.id">{{ stock.name }}</option>
                        </b-select>
                    </div>
                </div>
                <div class="mb-1" v-if="isErr">
                    <div class="control">
                        <p class="has-text-danger">Вы не выбрали склад!</p>
                    </div>
                </div>
            </div>
            <div class="card-content">
                <div class="field has-addons mb-1">
                    <div class="control">
                        <vue-simple-suggest
                            ref="wmsSuggestComponent"
                            v-model="searchString"
                            value-attribute="key"
                            :max-suggestions="0"
                            placeholder="Начните вводить наименование или артикул"
                            display-attribute="display"
                            :list="suggestionList"
                            :styles="{ defaultInput: {'input':true, 'is-medium':true} }"
                            :prevent-submit="true"
                            :debounce="200"
                            @select="onSuggestSelect">
                            <div slot="suggestion-item" slot-scope="{ suggestion, query }">
                                <div>
                                    <span>{{ suggestion.name }}</span>
                                </div>
                            </div>
                        </vue-simple-suggest>
                    </div>
                </div>
            </div>
        </div>
        <div class="card">
            <div class="card-content">
                <div class="content">
                    <table class="table">
                        <thead>
                        <th scope="col">Наименование</th>
                        <th scope="col">Штрихкод</th>
                        <th scope="col">Кол-во</th>
                        <th scope="col">Закупка</th>
                        <th scope="col">Удалить</th>
                        </thead>
                        <tbody>
                        <tr v-for="(item, key) in rowData" :key="key">
                            <td>{{ item.name }}</td>
                            <td>{{ item.uid }}</td>
                            <td>
                                <div class="quantity-toggle">
                                    <button @click="decrement(key)">&mdash;</button>
                                        <input type="text" :value="item.qty" readonly>
                                    <button @click="increment(key)">&#xff0b;</button>
                                </div>
                            </td>
                            <td>
                                <div class="posting-price">
                                    <input type="text" :value="item.price" @input="changePrice(key, $event)">
                                </div>
                            </td>
                            <td><font-awesome-icon @click="deleteItem(key)" class="has-text-danger" icon="trash" /></td>
                        </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
        <div class="card" v-if="!this.status">
            <div class="card-content">
                <p><a class="button is-primary is-rounded" @click="savePosting(false)">Сохранить</a>
                    &nbsp;&nbsp;&nbsp;
                    <a class="button is-success is-rounded" @click="savePosting(true)">Сохранить и провести</a></p>
            </div>
        </div>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex'
    import VueSimpleSuggest from 'vue-simple-suggest'
    import Loading from 'vue-loading-overlay'
    import Axios from 'axios'

    export default {
        name: "CreatePosting",
        computed: {
            ...mapGetters(
                {
                    stocks: 'transfer/stock',
                }),
        },
        async created() {
            this.$store.dispatch('transfer/getStocks')

            if (this.$route.params.postingId !== 'new' && this.$route.params.postingId !== undefined) {

                try {
                    const response = await Axios.get('/api/postings/' + this.$route.params.postingId)

                    if (response.status !== 200) {
                        throw 'posting not available'
                    }

                    this.status = response.data.status
                    this.rowData = response.data.products
                    this.document_id = response.data.id
                } catch (error) {
                    alert(error)
                }

            }
        },
        data: () => ({
            isErr: false,
            searchString: '',
            error: null,
            rowData: [],
            toStock: 1,
            suggestions: {},
            isLoading: false,
            fullPage: true,
            document_id: null,
            status: false
        }),
        methods: {
            async suggestionList(inputValue) {
                const {data} = await Axios.get(
                    '/api/wares/suggest',
                    {
                        params: {
                            text: inputValue
                        }
                    }
                )

                this.suggestions = data.success.map(
                    function (item, index) {
                        return {
                            key: index,
                            uid: item.barcode,
                            name: item.name,
                            display: item.name
                        }
                    }
                );
                return this.suggestions
            },
            onSuggestSelect: async function (suggest) {
                if (suggest) {
                    this.rowData.push({
                        uid:suggest.uid,
                        name:suggest.name,
                        qty:1
                    })
                    this.searchString = ''
                    this.$refs.wmsSuggestComponent.hideList()
                }
            },
            deleteItem: async function (key) {
                this.rowData.splice(key, 1)
            },
            increment (key) {
                this.rowData[key].qty++
            },
            decrement (key) {
                if(this.rowData[key].qty === 1) {
                    alert('Negative quantity not allowed')
                } else {
                    this.rowData[key].qty--
                }
            },
            savePosting: async function (status) {
                this.isLoading = true
                const {response} = await Axios.post(
                    '/api/postings',
                    {
                        products: this.rowData,
                        stock_id: this.stockId,
                        id: this.document_id,
                        status
                    }
                )
                this.isLoading = false
            },
            changePrice(key, event) {
                this.rowData[key].price = event.target.value
            }
        },
        components: {
            VueSimpleSuggest,
            Loading
        }
    }
</script>

<style scoped>

</style>