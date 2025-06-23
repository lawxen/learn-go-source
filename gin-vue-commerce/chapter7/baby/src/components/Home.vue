<template>
    <base-page :activation="activation"></base-page>

    <div class="category-con">
      <div class="category-banner">
        <div class="w1200">
          <img src="img/banner1.jpg">
        </div>
      </div>
    </div>
    <div class="floors">
      <div class="sk">
        <div class="sk_inner w1200">
          <div class="sk_hd">
            <a href="javascript:;">
              <img src="img/s_img1.jpg">
            </a>
          </div>
          <div class="sk_bd">
            <div class="layui-carousel" id="test1">
              <div carousel-item>
              <div class="item-box" v-for="(commodityInfo, k) in commodityInfos" :key="k">
                <div class="item" v-for="(c, key) in commodityInfo" :key="key">
                    <router-link :to="`/commodity/detail/${c.ID}`">
                        <img :src="path + c.img">
                    </router-link>
                    <div class="title">{{ c.name }}</div>
                    <div class="price">
                      <span>￥{{ c.discount }}</span>
                      <del>￥{{ c.price }}</del>
                    </div>
                </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="product-cont w1200" id="product-cont">
      <div class="product-item product-item1 layui-clear">
        <div class="left-title1">
          <h4><i>1F</i></h4>
          <img src="img/icon_gou.png">
          <h5>宝宝服饰</h5>
        </div>
        <div class="right-cont">
          <a href="javascript:;" class="top-img"><img src="img/img12.jpg" alt=""></a>
          <div class="img-box">
              <router-link v-for="(c, key) in clothes" :key="key" :to="`/commodity/detail/${c.ID}`">
                  <img :src="path + c.img">
              </router-link>
          </div>
        </div>
      </div>
      <div class="product-item product-item2 layui-clear">
        <div class="left-title2">
          <h4><i>2F</i></h4>
          <img src="img/icon_gou.png">
          <h5>奶粉辅食</h5>
        </div>
        <div class="right-cont">
          <a href="javascript:;" class="top-img"><img src="img/img12.jpg" alt=""></a>
          <div class="img-box">
              <router-link v-for="(c, key) in food"  :key="key" :to="`/commodity/detail/${c.ID}`">
                  <img :src="path + c.img">
              </router-link>
          </div>
        </div>
      </div>
      <div class="product-item product-item3 layui-clear">
        <div class="left-title3">
          <h4><i>3F</i></h4>
          <img src="img/icon_gou.png">
          <h5>宝宝用品</h5>
        </div>
        <div class="right-cont">
          <a href="javascript:;" class="top-img"><img src="img/img12.jpg"></a>
          <div class="img-box">
              <router-link v-for="(c, key) in goods" :key="key" :to="`/commodity/detail/${c.ID}`">
                  <img :src="path + c.img">
              </router-link>
          </div>
        </div>
      </div>
    </div>
    <footer-page></footer-page>
</template>

<script>
    export default {
        name: "Index",
        data() {
            return {
                activation: "",
                commodityInfos: [[], []],
                clothes:[],
                food: [],
                goods: [],
                path: this.$store.state.lookImgUrl

            }
        },
        mounted: function(){
            this.getcode(); //页面加载时自动执行
            // 执行layui的JS脚本
            window.layui.config({
                    base: 'js/'
                }).use(['mm', 'carousel'], function () {
                    var carousel = window.layui.carousel;
                    var option = {
                        elem: '#test1'
                        , width: '100%'
                        , arrow: 'always'
                        , height: '298'
                        , indicator: 'none'
                    }
                    carousel.render(option);
                });
        },
        methods: {
            getcode: function () {
                this.axios.get('/api/v1/home/'
                ).then(response => {
                    this.commodityInfos = response.data.data.commodityInfos
                    this.clothes = response.data.data.clothes
                    this.food = response.data.data.food
                    this.goods = response.data.data.goods
                    console.log(this.commodityInfos)
                })
                    .catch(function (error) {
                        console.log(error)
                    })
            },
        }
    }
</script>

<style scoped>

</style>
