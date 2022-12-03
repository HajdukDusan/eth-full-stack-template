<template>
    <div class="container-fluid">
        <h2>Content</h2>
        <h2>{{ test }}</h2>
    </div>
</template>

<script>
export default {
    name: 'Dashboard',

    data() {
        return {
            test: null
            
        };
    },

    created() {
        console.log(process.env);
        this.test = process.env.VUE_APP_TEST;
    },

    mounted() {
        
    },

    methods: {
        
    },
};
</script>

<style lang="scss" scoped>

</style>