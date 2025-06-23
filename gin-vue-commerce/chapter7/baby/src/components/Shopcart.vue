<template>
  <base-page :activation="activation"></base-page>
  <div class="banner-bg w1200">
    <h3>夏季清仓</h3>
    <p>宝宝被子、宝宝衣服3折起</p>
  </div>
  <div class="carts w1200">
    <div class="cart-table-th">
      <div class="th th-chk">
        <div class="select-all">
          <div class="cart-checkbox">
            <input class="check-all check" id="allCheckked" type="checkbox" value="true">
          </div>
          <label>&nbsp;&nbsp;全选</label>
        </div>
      </div>
      <div class="th th-item">
        <div class="th-inner">商品</div>
      </div>
      <div class="th th-price">
        <div class="th-inner">单价</div>
      </div>
      <div class="th th-amount">
        <div class="th-inner">数量</div>
      </div>
      <div class="th th-sum">
        <div class="th-inner">小计</div>
      </div>
      <div class="th th-op">
        <div class="th-inner">操作</div>
      </div>
    </div>
    <div class="OrderList">
      <div class="order-content" id="list-cont">

        <ul class="item-content layui-clear" v-for="(c, key) in cartInfos" :key="key">
          <li class="th th-chk">
            <div class="select-all">
              <div class="cart-checkbox">
                <input class="CheckBoxShop check" type="checkbox" num="all"
                       name="select-all" value="true">
              </div>
            </div>
          </li>
          <li class="th th-item">
            <div class="item-cont">
              <a href="javascript:;"><img :src="path + c.Commodities.img"></a>
              <div class="text">
                <div class="title">{{ c.Commodities.name }}</div>
                <p><span>{{ c.Commodities.sezes }}</span></p>
              </div>
            </div>
          </li>
          <li class="th th-price">
            <span class="th-su">{{ c.Commodities.price }}</span>
          </li>
          <li class="th th-amount">
            <div class="box-btn layui-clear">
              <div class="less layui-btn">-</div>
              <input class="Quantity-input" :value="c.quantity" disabled="disabled">
              <div class="add layui-btn">+</div>
            </div>
          </li>
          <li class="th th-sum">
            <span class="sum">0</span>
          </li>
          <li class="th th-op" @click="delCarInfo(c.ID)">
            <span>删除</span>
            <p hidden="hidden" class="cartPrimary">{{ c.ID }}</p>
          </li>
        </ul>

      </div>
    </div>
    <div class="FloatBarHolder layui-clear">
      <div class="th th-chk">
        <div class="select-all">
          <div class="cart-checkbox">
            <input class="check-all check" name="select-all" type="checkbox" value="true">
          </div>
          <label>&nbsp;&nbsp;已选<span class="Selected-pieces">0</span>件</label>
        </div>
      </div>
      <div class="th batch-deletion" @click="delCarInfo(username)">
        <span class="batch-dele-btn">删除全部</span>
        <p hidden="hidden" id="userId">{{ username }}</p>
      </div>
      <div class="th Settlement">
        <button class="layui-btn" @click="pays">结算</button>
      </div>
      <div class="th total">
        <p>应付：<span class="pieces-total">￥0.00</span></p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Shopcart",
  data() {
    return {
      activation: 'shopcart',
      cartInfos: [],
      username: this.$store.state.username,
      path: this.$store.state.lookImgUrl,
      token: this.$store.state.token,
    }
  },
  beforeMount: function () {
    this.getcode(); //页面加载时自动执行
  },
  updated: function () {
    this.getSum(); //页面更新完成执行JS计算价钱
  },
  methods: {
    getcode() {
      this.axios({
        method: 'get',
        url: '/api/v1/shopper/shopcart/',
        headers: {
          'Authorization': this.token
        }
      }).then(response => {
        this.cartInfos = response.data.data
        this.getSum()
      })
          .catch(function (error) {
            console.log(error)
          })
    },
    pays: function () {
      var cartID = []
      var uls = window.document.getElementById('list-cont').getElementsByTagName('ul')
      for (var i = 0; i < uls.length; i++) {
        if (uls[i].getElementsByTagName('input')[0].checked) {
          var carts = window.document.getElementsByClassName('cartPrimary')[i].innerHTML
          cartID.push(carts)
        }
      }
      var d = {
        'total': window.layui.$(".pieces-total").text(),
        'cartId': cartID
      }
      this.axios({
        method: 'post',
        url: '/api/v1/shopper/pays/',
        data: d,
        headers: {
          'Authorization': this.token
        }
      }).then(response => {
        if (response.data.data !== '') {
          window.location.href = response.data.data
        }
      })
          .catch(function (error) {
            console.log(error)
          })
    },
    delCarInfo(style) {
      var d
      if (style === this.username) {
        d = {'cartId': 0}
      } else {
        d = {'cartId': style}
      }
      this.axios({
        method: 'post',
        url: '/api/v1/shopper/delete/',
        data: d,
        headers: {
          'Authorization': this.token
        }
      }).then(response => {
        if (response.data.state === 'success') {
          // 删除成功
          this.getcode(); //重新加载购物车
        }
      })
          .catch(function (error) {
            console.log(error)
          })
    },
    getSum: function () {
      window.layui.config({
        base: 'js/'
      }).use(['jquery', 'element', 'car'], function () {
        var $ = window.layui.$;
        var car = window.layui.car;
        car.init();
        $(function () {
          var counts = 0;
          $(".sum").each(function (i, e) {
            var quantity = $('.th-su')[i].innerHTML
            var price = $('.Quantity-input')[i].value
            e.innerHTML = (quantity * price).toFixed(2)
            counts = counts * 1 + e.innerHTML * 1
          });
          $(".pieces-total").text("￥" + counts.toFixed(2))
        });
      });
    }
  },
}
</script>

<style scoped>

</style>
