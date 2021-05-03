<template>
    <div class="content">
        <div class="card">
            <div class="card-content">
                <div class="field has-addons mb-1">
                    <div class="control">
                    <b-field label="Документ-основание">
                        <b-input type="text" placeholder="Введите номер документа"  v-model="baseDocumentId"></b-input>
                        <p class="control">
                            <b-button class="button is-primary" @click="setBaseDocumentId">Заполнить по документу-основанию</b-button>
                        </p>
                    </b-field>
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
                        <th scope="col">Стеллаж</th>
                        <th scope="col">Удалить</th>
                        </thead>
                        <tbody>
                        <tr v-for="(item, key) in rowData" :key="key">
                            <td>{{ item.name }}</td>
                            <td>{{ item.qty }}</td>
                            <td>
                                <div>
                                    <input type="text" @input="setCell(item.id, $event.target.value)" :value="item.cell_id" style="width:50px">
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
                </p>
            </div>
        </div>
    </div>
</template>

<script>
    import Axios from 'axios';

    Axios.defaults.baseURL = 'http://localhost:8010';

    export default {
        name: "AllocationEdit",
        async created() {
            console.log(this.$route.params.allocationId);
            if (this.$route.params.allocationId !== 'new' && this.$route.params.allocationId !== undefined) {

                try {
                    const response = await Axios.get('/allocation/' + this.$route.params.allocationId)

                    if (response.status !== 200) {
                        throw 'receipt not available'
                    }

                    this.status = response.data.status
                    this.rowData = response.data.products
                    this.documentId = response.data.id
                    console.log(this.documentId);
                } catch (error) {
                    console.log(error)
                }

            }
        },
        data: () => ({
            isErr: false,
            error: null,
            rowData: [],
            fullPage: true,
            documentId: null,
            status: false,
            cells: []
        }),
        computed: {
            baseDocumentId: {
                get () { return this.$store.getters['allocation/documentReceiptId'] },
                set (value) { this.$store.dispatch('allocation/setDocumentId', value) }
            },
        },
        methods: {
            setBaseDocumentId: async function () {
                if (this.baseDocumentId !== '') {
                    try {
                        const response = await Axios.get('/receipt/' + this.baseDocumentId);
                        await Axios.put('/allocation/' + this.documentId, {
                            document_id: Number(this.documentId),
                            document_type: "receipt"
                        });
                        for (let key in response.data.products) {
                            let row = new Object();
                            console.log(response.data.products[key]);
                            row.id = response.data.products[key].id;
                            row.name = response.data.products[key].name;
                            row.product_id = response.data.products[key].product_id;
                            row.qty = response.data.products[key].qty;
                            row.cell = '';
                            this.rowData.push(row);
                            await Axios.post('/allocationrow', {
                                allocation_id: this.documentId,
                                product_id: row.product_id,
                                qty: row.qty,
                                cell_id: 0
                            });
                        }
                        // response.data.qty = 1; 
                        // this.barcode = '';
                        // this.$refs.barcode.focus();
                        // const responseRow = await Axios.post('/receiptrow', {
                        //     receipt_id: this.documentId,
                        //     product_id: response.data.id,
                        //     qty: response.data.qty
                        // });
                        // response.data.id = responseRow.data.id;
                        // this.rowData.push(response.data);
                        // console.log(this.rowData);
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
                await Axios.patch('/allocation/' + this.$route.params.allocationId,
                    {
                        status: this.status == 2 ? 1 : 2
                    }
                )
                this.status = this.status == 2 ? 1 : 2
            },
            createAllocation() {
                this.$store.dispatch('allocation/setDocumentId', this.documentId);
                this.$router.push('/allocation/new');
            },
            setCell: async function (id, key) {
                console.log(id,key);
                await Axios.put('/allocationrow/' + id,
                    {
                        cell_id: Number(key)
                    }
                )
            },
        },
        components: {
            // VueSimpleSuggest,
            // Loading
        }
    }
</script>

<style scoped>

</style>