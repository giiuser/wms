<template>
    <div class="content">
        <div class="card">
            <div class="card-content">
                <div class="field has-addons mb-1">
                    <div class="control">
                        <input class="input" type="text" :ref="'barcode'" @input="setBarcode" placeholder="Сканируйте штрихкод"  v-model="barcode">
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
                        <th scope="col">Кол-во</th>
                        <th scope="col">Удалить</th>
                        </thead>
                        <tbody>
                        <tr v-for="(item, key) in rowData" :key="key">
                            <td>{{ item.name }}</td>
                            <td>
                                <div class="quantity-toggle">
                                    <input type="text" :value="item.qty" readonly style="width:50px">
                                </div>
                            </td>
                            <td><font-awesome-icon @click="deleteRow(item.id, key)" class="has-text-danger" icon="trash" /></td>
                        </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
        <div class="card">
            <div class="card-content">
                <p>
                    <a class="button is-success is-rounded" @click="posting()">{{ this.status == 2 ? 'Распровести' : 'Провести' }}</a>
                    &nbsp;&nbsp;&nbsp;
                    <a class="button is-primary is-rounded" v-if="this.status == 2" @click="createAllocation()">Создать размещение</a>
                </p>
            </div>
        </div>
    </div>
</template>

<script>
    import Axios from 'axios';

    Axios.defaults.baseURL = 'http://localhost:8010';

    export default {
        name: "ReceiptEdit",
        async created() {
            if (this.$route.params.receiptId !== 'new' && this.$route.params.receiptId !== undefined) {

                try {
                    const response = await Axios.get('/receipt/' + this.$route.params.receiptId)

                    if (response.status !== 200) {
                        throw 'receipt not available'
                    }

                    this.status = response.data.status
                    this.rowData = response.data.products
                    this.documentId = response.data.id
                } catch (error) {
                    console.log(error)
                }

            }
        },
        data: () => ({
            barcode: '',
            isErr: false,
            searchString: '',
            error: null,
            rowData: [],
            toStock: 1,
            suggestions: {},
            isLoading: false,
            fullPage: true,
            documentId: null,
            status: false
        }),
        methods: {
            setBarcode: async function () {
                if (this.barcode !== '') {
                    try {
                        console.log(this.barcode);
                        const response = await Axios.get('/product/' + this.barcode);
                        response.data.qty = 1; 
                        this.barcode = '';
                        this.$refs.barcode.focus();
                        const responseRow = await Axios.post('/receiptrow', {
                            receipt_id: this.documentId,
                            product_id: response.data.id,
                            qty: response.data.qty
                        });
                        response.data.id = responseRow.data.id;
                        this.rowData.push(response.data);
                        console.log(this.rowData);
                    } catch (error) {
                        console.log(error);
                    }
                } else {
                    console.log('error')
                }
            },
            deleteRow: async function (id, key) {
                await Axios.delete('/receiptrow/' + id);
                this.rowData.splice(key, 1);
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
            posting: async function () {
                await Axios.patch('/receipt/' + this.$route.params.receiptId,
                    {
                        status: this.status == 2 ? 1 : 2
                    }
                )
                this.status = this.status == 2 ? 1 : 2
            },
            createAllocation() {
                this.$store.dispatch('allocation/setDocumentId', this.documentId);
                this.$router.push('/allocation/new').catch(() => {});
            }
        },
        components: {

        }
    }
</script>

<style scoped>

</style>