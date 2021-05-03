<template>
    <div class="wares-table">
        <section>
            <b-field label="Введите наименование">
                <b-input v-model="name"></b-input>
            </b-field>

            <b-field v-if="errors.length">
                <ul>
                    <li v-for="(error, key) in errors" :key="key">{{ error }}</li>
                </ul>
            </b-field>

            <b-field>
                <a class="button is-success" @click="save">Сохранить</a>
            </b-field>
        </section>
    </div>
</template>

<script>
import Axios from "axios";

Axios.defaults.baseURL = 'http://localhost:8010';

export default {
    name: "ProductEdit",
    data: () => ({
        name: "",
        errors: []
    }),
    async created() {
        if (this.$route.params.productId !== 'new' && this.$route.params.productId !== undefined) {

            try {
                const response = await Axios.get('/product/' + this.$route.params.productId)

                if (response.status !== 200) {
                    throw 'product not available'
                }

                this.name = response.data.name
            } catch (error) {
                alert(error)
            }

            }
    },
    methods: {
        save: function() {
            if (this.checkForm() === true) {
                Axios.put('/product/' + this.$route.params.productId, {
                    name: this.name
                })
                    .then(response => {
                        console.log(response)
                        this.$router.push("/products");
                    })
                    .catch(error => {
                        console.log(error);
                    });
            }
        },
        checkForm: function() {
            if (this.name) {
                return true;
            }

            this.errors = [];

            if (!this.name) {
                this.errors.push("Требуется указать имя.");
            }
        }
    },
    mounted() {
        
    }
};
</script>
