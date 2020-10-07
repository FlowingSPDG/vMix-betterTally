<script>
export default {
  name: "api",
  methods: {
    async GetInputs() {
      try {
        const res = await this.axios.get("/api/inputs");
        return res.data;
      } catch (err) {
        throw new Error(err);
      }
    },
  }
};
</script>