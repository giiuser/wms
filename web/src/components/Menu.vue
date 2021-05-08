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
            <b-navbar-dropdown label="Склад">
                <b-navbar-item href="/products">
                    Товары
                </b-navbar-item>
                <b-navbar-item href="/cells">
                    Полки
                </b-navbar-item>
                <b-navbar-item href="/stocks">
                    Остатки
                </b-navbar-item>
            </b-navbar-dropdown>
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
        
        </template>
    </b-navbar>
</template>