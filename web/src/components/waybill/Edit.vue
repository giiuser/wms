<template>
    <div class="content">
        <div class="card">
            <div class="card-content">
                <div class="field has-addons mb-1">
                    <div class="control">
                    <b-field label="Документ-основание">
                        <b-input type="text" placeholder="Введите номер документа"  v-model="baseDocumentId" :disabled="this.status === 2"></b-input>
                        <p class="control">
                            <b-button class="button is-primary" @click="setBaseDocumentId" :disabled="this.status === 2">Заполнить по документу-основанию</b-button>
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
                        <th scope="col">Производитель</th>
                        <th scope="col">Кол-во</th>
                          </thead>
                        <tbody>
                        <tr v-for="(item, key) in rowData" :key="key">
                            <td>{{ item.name }}</td>
                            <td>{{ item.brand }}</td>
                            <td>{{ item.qty }}</td>
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

    // Axios.defaults.baseURL = 'http://localhost:8010';

    export default {
        name: "WaybillEdit",
        async created() {
            console.log(this.$route.params.allocationId);
            if (this.$route.params.waybillId !== 'new' && this.$route.params.waybillId !== undefined) {

                try {
                    const response = await Axios.get('/waybill/' + this.$route.params.waybillId)

                    if (response.status !== 200) {
                        throw 'waybill not available';
                    }

                    this.status = response.data.status;
                    this.rowData = response.data.products;
                    this.documentId = response.data.id;
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
                get () { return this.$store.getters['waybill/documentReceiptId'] },
                set (value) { this.$store.dispatch('waybill/setDocumentId', value) }
            },
        },
        methods: {
            setBaseDocumentId: async function () {
                if (this.baseDocumentId !== '') {
                    try {
                        await Axios.delete('/waybillrows/' + this.documentId);
                        this.rowData = [];
                        const response = await Axios.get('/collect/' + this.baseDocumentId);
                        await Axios.put('/waybill/' + this.documentId, {
                            document_id: Number(this.documentId),
                            document_type: "collect"
                        });
                        for (let key in response.data.products) {
                            let row = new Object();
                            console.log(response.data.products[key]);
                            row.id = response.data.products[key].id;
                            row.name = response.data.products[key].name;
                            row.brand = response.data.products[key].brand;
                            row.product_id = response.data.products[key].product_id;
                            row.qty = response.data.products[key].qty;
                            row.cell = '';
                            this.rowData.push(row);
                            await Axios.post('/waybillrow', {
                                waybill_id: this.documentId,
                                product_id: row.product_id,
                                qty: row.qty
                            });
                        }
                    } catch (error) {
                        console.log(error);
                    }
                } else {
                    console.log('error')
                }
            },
            deleteRow: async function (id, key) {
                await Axios.delete('/waybillrow/' + id);
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
                await Axios.patch('/waybill/' + this.$route.params.waybillId,
                    {
                        status: this.status == 2 ? 1 : 2
                    }
                )
                this.status = this.status == 2 ? 1 : 2
            },
            createAllocation() {
                this.$store.dispatch('waybill/setDocumentId', this.documentId);
                this.$router.push('/waybill/new');
            },
            setCell: async function (id, key) {
                console.log(id,key);
                await Axios.put('/waybillrow/' + id,
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