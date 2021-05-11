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
                <a class="button is-success" @click="create">Создать</a>
            </b-field>
        </section>
    </div>
</template>

<script>
import Axios from "axios";

// Axios.defaults.baseURL = 'http://localhost:8010';

export default {
    name: "CellCreate",
    data: () => ({
        name: "",
        brand: "",
        errors: []
    }),
    methods: {
        loadAsyncData: function() {

        },
        create: function() {
            if (this.checkForm() === true) {
                Axios.post("/cell", {
                    name: this.name
                })
                    .then(response => {
                        console.log(response)
                        this.$router.push("/cells");
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
