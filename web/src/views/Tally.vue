<template>
    <div v-loading.fullscreen.lock="loading" :style="`background-color:currentColor;color:${backgroundColor};display:inline-block;width:100vw;height:100vh;`"></div>
</template>


<script>
// @ is an alias to /src

export default {
  name: "TALLY",
  components: {
  },
  data() {
    return {
      loading:true,
      ws:null,
      status:"Off",
      backgroundColor:"grey"
    };
  },
  async mounted() {
      const key = this.$route.query.key

      const url = `ws://${location.hostname}:${location.port}/api/ws`
      this.ws = new WebSocket(url);
      console.log(url)
      // websocketをオープンした時
      this.ws.onopen = (event) => {
        console.log("### websocket.onopen()");
        console.log(event)
        this.loading=false
      };

      // websocketでメッセージを受信した時
      this.ws.onmessage = (event) => {
        console.log("### websocket.onmessage()");
        const obj = JSON.parse(event.data)
        console.log(obj)
        this.loading=false
        this.getMyStatus(obj)
      };

      // websocketでエラーが発生した時
      this.ws.onerror = (event) => {
        console.log("### websocket.onerror()");
        console.log(event);
        this.loading = true
      };

      // websocketをクローズした時
      this.ws.onclose = (event) => {
        console.log("### websocket.onclose()");
        console.log(event);
        this.loading = true
      };
  },
  watch: {
    status: function (newStatus, oldStatus) {
      switch(newStatus){
        case "Off":
            this.backgroundColor = "grey"
            break
        case "Preview":
            this.backgroundColor = "green"
            break
        case "Program":
            this.backgroundColor = "red"
            break
      }
    }
  },
  methods: {
    getMyStatus: function(obj){
        for(const k in obj) {
            console.log(`k : ${k}`)
            console.log(`this.$route.query.key : ${this.$route.query.key}`)
           if(this.$route.query.key == k){
               console.log(`this.$route.query.key == k : ${this.$route.query.key == k}`)
               this.status = this.getStatus(obj[k])
               break
           }
        }
    },
    isPreview:function(status) {
        return status == 2
    },
    isProgram:function(status){
        return status == 1
    },
    isOff:function(status){
        return status == 0
    },
    getStatus:function(status){
        console.log(`getStatus : ${status}`)
        switch(status){
            case 0:
                return "Off"
            case 1:
                return "Program"
            case 2:
                return "Preview"
            default:
                return "Unknown"
        }
    }
  }
};
</script>

<style>

</style>