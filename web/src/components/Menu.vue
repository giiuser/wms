<template>
    <b-navbar>
        <template #brand>
            <b-navbar-item tag="router-link" :to="{ path: '/' }">
                <img
                    src="https://raw.githubusercontent.com/buefy/buefy/dev/static/img/buefy-logo.png"
                    alt="Lightweight UI components for Vue.js based on Bulma"
                >
            </b-navbar-item>
        </template>
        <template #start>
            <b-navbar-item href="/">
                Домой
            </b-navbar-item>
            <b-navbar-item href="/products">
                Товары
            </b-navbar-item>
            <b-navbar-dropdown label="Документы">
                <b-navbar-item href="/receipts">
                    Оприходования
                </b-navbar-item>
                <b-navbar-item href="/allocations">
                    Размещения
                </b-navbar-item>
                <b-navbar-item href="/collects">
                    Заказы на сборку
                </b-navbar-item>
                <b-navbar-item href="/waybills">
                    Расходные накладные
                </b-navbar-item>
            </b-navbar-dropdown>
        </template>

        <template #end>
            <b-navbar-item tag="div">
                <div class="buttons">
                    <a class="button is-primary">
                        <strong>Sign up</strong>
                    </a>
                    <a class="button is-light">
                        Log in
                    </a>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>