<template>
  <div class="home">
    <el-table ref="singleTable" v-if="Array.isArray(inputs)" :default-sort = "{prop: 'Number', order: 'ascending'}" :data="inputs" style="width:85%;margin:auto;" v-loading="loading" highlight-current-row>
      <el-table-column label="Number" prop="Number" sortable > </el-table-column>
      <el-table-column label="Name" prop="Name" sortable> </el-table-column>
      <el-table-column label="vMix functions" width="300">
        <template slot-scope="scope" style="">
          <el-button-group>
            <el-button plain size="small" round icon="el-icon-copy-document" @click="openTally(scope.row.Key)">TALLY</el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "Home",
  components: {
  },
  data() {
    return {
      loading:false,
      inputs:[]
    };
  },
  async mounted() {
    this.inputs = await this.GetInputs()
    console.log(this.inputs)
  },
  methods: {
    async openTally(key){
      window.open(`/#/tally?key=${key}`, '_blank')
    }
  }
};
</script>
