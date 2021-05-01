<template>
<div>
        <v-toolbar rounded tile>
        <v-app-bar-nav-icon> </v-app-bar-nav-icon>
        <v-app-bar-title>
            WMS by Sergey Rukhmanov
        </v-app-bar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items>
            <v-btn flat to="/">
                Домой
            </v-btn>
            <v-btn flat to="/products">
                Товары
            </v-btn>
            <v-menu :rounded="rounded" open-on-hover offset-y transition="slide-x-transition" bottom right>
                <template v-slot:activator="{ on, attrs }">
                    <v-btn flat v-bind="attrs" v-on="on">
                        Документы
                    </v-btn>
                </template>
                <v-list dense>
                    <v-list-item v-for="(item, index) in documents" :key="index" router :to="item.link">
                       <v-list-item-action>
                            <v-list-item-title>{{ item.title }}</v-list-item-title>
                        </v-list-item-action>
                    </v-list-item>
                </v-list>
            </v-menu>
        </v-toolbar-items>
    </v-toolbar>
</div>
</template>

<script>
export default {
    data: () => ({
        activate: true,
        theme: 1,
        themes: [{
                text: "Dark",
                icon: "mdi-clock"
            },
            {
                text: "Light",
                icon: "mdi-account"
            }
        ],
        mini: true,
        documents: [{
                icon: "mdi-domain",
                title: "Оприходования",
                link: "/receipts"
            },
            {
                icon: "mdi-message-text",
                title: "Размещение товаров",
                link: "/allocations"
            },
            {
                icon: "mdi-domain",
                title: "Заказы на сборку",
                link: "/collects"
            },
            {
                icon: "mdi-message-text",
                title: "Расходные накладные",
                link: "/waybills"
            }
        ]
    })
};
</script>