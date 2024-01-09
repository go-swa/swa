<template>
  <div>
    <textarea
      ref="mycode"
      v-model="value"
      class="codesql"
    />
  </div>
</template>

<script>
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/blackboard.css'
import 'codemirror/addon/hint/show-hint.css'
const CodeMirror = require('codemirror/lib/codemirror')
require('codemirror/addon/edit/matchbrackets')
require('codemirror/addon/selection/active-line')
require('codemirror/mode/sql/sql')
require('codemirror/addon/hint/show-hint')
require('codemirror/addon/hint/sql-hint')

export default {
  name: 'SqlEditor',
  props: {
    value: {
      type: String,
      default: ''
    },
    sqlStyle: {
      type: String,
      default: 'default'
    },
    readOnly: {
      type: [Boolean, String]
    }
  },
  data() {
    return {
      editor: null
    }
  },
  computed: {
    newVal() {
      if (this.editor) {
        return this.editor.getValue()
      }
    }
  },
  watch: {
    newVal(newV, oldV) {
      if (this.editor) {
        this.$emit('changeTextarea', this.editor.getValue())
      }
    }
  },
  mounted() {
    const mime = 'text/x-mariadb'
    const theme = 'blackboard'
    this.editor = CodeMirror.fromTextArea(this.$refs.mycode, {
      value: this.value,
      mode: mime,
      indentWithTabs: true,
      smartIndent: true,
      lineNumbers: true,
      matchBrackets: true,
      cursorHeight: 1,
      lineWrapping: true,
      readOnly: this.readOnly,
      theme: theme,
      autofocus: true,
      extraKeys: { 'Ctrl': 'autocomplete' },
      hintOptions: {
        completeSingle: false
      }
    })
    this.editor.on('inputRead', () => {
      this.editor.showHint()
    })
  },
  methods: {
    setVal() {
      if (this.editor) {
        if (this.value === '') {
          this.editor.setValue('')
        } else {
          this.editor.setValue(this.value)
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.CodeMirror {
  border: 1px solid black;
  font-size: 13px;
}

.CodeMirror-hints{
  z-index: 9999 !important;
}
</style>
