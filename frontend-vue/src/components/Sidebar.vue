<template>
    <div>
        <h4 class="sidebar-header pt-3">Sidebar</h4>
        <hr>
        <ul class="nav nav-pills flex-column mb-auto px-3">
            <li v-for="item in items" :key="item.name" class="nav-item">
                <span @click="active = item.name" class="nav-link" :class="{ active: active == item.name }">
                    <i :class="item.icon" class="mx-3"></i>{{ item.name }}
                    </span>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    name: 'Sidebar',

    data() {
        return {
            active: "Dashboard",
            items: [
                {
                    name: "Dashboard",
                    to: '/',
                    icon: 'fa fa-home'
                },
                {
                    name: "Wallet",
                    to: '/wallet',
                    icon: 'fa fa-wallet'
                }
            ]
            
        };
    },

    mounted() {
        
    },

    methods: {
        
    },
};
</script>

<style scoped>

.sidebar-header {
    text-align: center;
}

span {
    transition: background-color 0.5s;
    cursor: pointer;
}

</style>