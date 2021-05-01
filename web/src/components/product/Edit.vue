<template>
    <div class="wares-table">
        <section>
            <b-field label="Введите наименование">
                <b-input v-model="name"></b-input>
            </b-field>

            <b-field label="Введите артикул">
                <b-input v-model="article"></b-input>
            </b-field>

            <b-field v-if="errors.length">
                <ul>
                    <li v-for="(error, key) in errors" :key="key">{{ error }}</li>
                </ul>
            </b-field>

            <b-field>
                <a class="button is-success" @click="newWare">Создать</a>
            </b-field>
        </section>
    </div>
</template>

<script>
import Axios from "axios";

Axios.defaults.baseURL = 'http://localhost:8010';

export default {
    name: "CreateWare",
    data: () => ({
        name: "",
        article: "",
        color: 0,
        size: 0,
        colors: [],
        sizes: [],
        errors: []
    }),
    methods: {
        loadAsyncData: function() {

        },
        newWare: function() {
            if (this.checkForm() === true) {
                Axios.post("/api/wares", {
                    name: this.name,
                    article: this.article,
                    color: this.color,
                    size: this.size
                })
                    .then(response => {
                        this.$router.push("/wares");
                    })
                    .catch(error => {
                        console.log(error);
                    });
            }
        },
        checkForm: function() {
            if (this.name && this.article && this.color && this.size) {
                return true;
            }

            this.errors = [];

            if (!this.name) {
                this.errors.push("Требуется указать имя.");
            }
            if (!this.article) {
                this.errors.push("Требуется указать артикул.");
            }
            if (!this.color) {
                this.errors.push("Требуется указать цвет.");
            }
            if (!this.size) {
                this.errors.push("Требуется указать размер.");
            }
        }
    },
    mounted() {
        this.loadAsyncData();
    }
};
</script>
