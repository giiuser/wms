<template>
    <div class="content">
        <div class="card">
            <div class="card-content">
                <div class="field has-addons mb-1">
                    <div class="control">
                        <input class="input" type="text" :ref="'barcode'" @input="setBarcode" placeholder="Сканируйте штрихкод"  v-model="barcode" :disabled="this.status === 2">
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
                        <th scope="col">Полка</th>
                        <th scope="col">Удалить</th>
                        </thead>
                        <tbody>
                        <tr v-for="(item, key) in rowData" :key="key">
                            <td>{{ item.name }}</td>
                            <td>{{ item.brand }}</td>
                            <td>{{ item.qty }}</td>
                            <td>
                                <div>
                                    <input type="text" @input="setCell(item.id, $event.target.value)" :value="item.cell_name.String" style="width:50px" :disabled="status === 2">
                                </div>
                            </td>
                            <td><font-awesome-icon v-if="status !== 2" @click="deleteRow(item.id, key)" class="has-text-danger" icon="trash" /></td>
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
                    <a class="button is-primary is-rounded" v-if="this.status == 2" @click="createWaybill()">Создать накладную на отгрузку</a>
                </p>
            </div>
        </div>
    </div>
</template>

<script>
    import Axios from 'axios';

    // Axios.defaults.baseURL = 'http://localhost:8010';

    export default {
        name: "CollectEdit",
        async created() {
            if (this.$route.params.collectId !== 'new' && this.$route.params.collectId !== undefined) {

                try {
                    const response = await Axios.get('/collect/' + this.$route.params.collectId)

                    if (response.status !== 200) {
                        throw 'collect not available'
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
                        const responseRow = await Axios.post('/collectrow', {
                            collect_id: this.documentId,
                            product_id: response.data.id,
                            cell_id: 0,
                            qty: response.data.qty
                        });
                        response.data.id = responseRow.data.id;
                        response.data.cell_id = 0;
                        response.data.cell_name = '';
                        console.log(response.data);
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
                await Axios.delete('/collectrow/' + id);
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
                await Axios.patch('/collect/' + this.$route.params.collectId,
                    {
                        status: this.status == 2 ? 1 : 2
                    }
                )
                this.status = this.status == 2 ? 1 : 2
            },
            createWaybill() {
                this.$store.dispatch('waybill/setDocumentId', this.documentId);
                this.$router.push('/waybill/new').catch(() => {});
            },
            setCell: async function (id, key) {
                console.log(id,key);
                await Axios.put('/collectrow/' + id,
                    {
                        cell_id: Number(key)
                    }
                )
            },
        },
        components: {

        }
    }
</script>

<style scoped>

</style>