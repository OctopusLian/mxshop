<template>
  <div class="banner-warp w-max ct bgwh mb30 ovh border-eee" >
    <swiper :options="swiperOption">
      <swiper-slide v-for="item in banners" :key="item.goods">
        <router-link :to="'/app/home/productDetail/'+item.url" target = _blank> <img :src="item.image" width="1200" height="500" alt="" /></router-link>
      </swiper-slide>
      <div class="swiper-pagination" slot="pagination"></div>
    </swiper>
  </div>
</template>
<style>
  .banner-warp{
    height:300px;
  }
</style>



<script>
  import { swiper, swiperSlide } from 'vue-awesome-swiper'
  import {bannerGoods} from '../../api/api'

  export default {
    components: {
      swiper,
      swiperSlide,
    },
    data() {

      return {

        swiperOption: {
          pagination: '.swiper-pagination',
          paginationClickable: true,
          autoplay: 2500,
          autoplayDisableOnInteraction: false,
        },
        banners:[]

      }

    },
    methods:{
      getBanner(){
        bannerGoods()
          .then((response)=> {
            console.log(response)
            //跳转到首页页response.body面
            this.banners = response.data
          })
          .catch(function (error) {
            console.log(error);
          });
      }
    },
    created(){
      this.getBanner();
    }

  }

</script>
