<template>
  <div class="table-responsive">
    <table>
      <thead>
        <tr v-for="(tr, index) in combineHeads" :key="index">
          <th v-for="cell in tr" :key="cell.__index" :rowspan="cell.rowspan" :colspan="cell.colspan">
            <div
              :class="{ 'col-corner-bg': cell.isCorner }"
              :style="{ 'min-height': _getMinHeightByRowCount(cell.rowspan) }"
            >
              {{ cell.isCorner ? (rowPaths.length + ' x ' + colPaths.length) : cell.value }}
            </div>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(tr, index) in combineValues" :key="index">
          <th
            v-for="cell in trhead"
            :key="cell.__index"
            :rowspan="cell.rowspan"
            :colspan="cell.colspan"
          >
            <div :style="{ 'min-height': _getMinHeightByRowCount(cell.rowspan) }">
              {{ cell.value }}
            </div>
          </th>
          <td v-for="cell in tr.data" :key="cell.__index" :rowspan="cell.rowspan" :colspan="cell.colspan">
            <div :style="{ 'min-height': _getMinHeightByRowCount(cell.rowspan) }">
              {{ cell.value }}
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { mergeBaseInfo, convertPathToMap, getHeightByCount, SEPARATOR } from '@/utils/visual-chart'

export default {
  name: 'PivotTable',
  props: {
    data: {
      type: Array,
      default: () => []
    },
    rows: {
      type: Array,
      default: () => []
    },
    columns: {
      type: Array,
      default: () => []
    },
    values: {
      type: Array,
      default: () => []
    }
  },
  data: () => ({
    localRows: [],
    localColumns: [],
    localValues: [],
    localData: [],
    calcData: [],
    Separator: SEPARATOR,
    combineHeads: [],
    combineValues: []
  }),
  computed: {
    watchAllProps() {
      const { rows, columns, values, data } = this
      return { rows, columns, values, data }
    },
    rowPaths() {
      const _paths = this._combineRowPaths(
        this.localData,
        ...this.localRows.map(({ key, values }) => {
          return { key, values }
        })
      )
      return _paths
    },
    colPaths() {
      const keys = this.localColumns.map(({ values }) => values)
      if (this.localValues.length) {
        keys.push(this.localValues.map(({ key }) => key))
      }
      const _paths = this._combineColPaths(...keys)
      return _paths
    },
    colHeads() {
      const _rows = this.localColumns.map(() => [])
      const valuesLen = this.localValues.length
      if (valuesLen) {
        _rows.push([])
      }
      const colSpans = {}
      this.colPaths.forEach((path, pathIndex) => {
        const pathValues = path.split(this.Separator)
        const currPath = []
        _rows.forEach((row, rowIndex) => {
          const cellData = {}
          const currVal = pathValues[rowIndex] || ''
          const isLastRow = rowIndex === _rows.length - 1
          currPath.push(currVal)
          const baseX = rowIndex
          const baseY = this.localRows.length + pathIndex
          if (!isLastRow) {
            let compareVal = valuesLen
            for (let i = rowIndex; i < this.localColumns.length - 1; i++) {
              compareVal *= this.localColumns[rowIndex + 1].values.length
            }
            const currColSpan = colSpans[rowIndex] || {}
            const currColSpanVal = (currColSpan[currPath.join(this.Separator)] || 0) + 1
            currColSpan[currPath.join(this.Separator)] = currColSpanVal
            colSpans[rowIndex] = currColSpan
            if (currColSpanVal === 1) {
              row.push(
                Object.assign(
                  cellData,
                  mergeBaseInfo({
                    __index: `${baseX}-${baseY}`,
                    x: baseX,
                    y: baseY,
                    colspan: compareVal,
                    path: currPath.filter((item) => !!item),
                    value: currVal
                  })
                )
              )
            }
          } else {
            row.push(
              Object.assign(
                cellData,
                mergeBaseInfo({
                  __index: `${baseX}-${baseY}`,
                  x: baseX,
                  y: baseY,
                  path: currPath.filter((item) => !!item),
                  value: this.localValues.find(({ key }) => key === currVal).label
                })
              )
            )
          }
        })
      })
      return _rows
    },
    rowHeads() {
      const _columns = []
      const columnsLen = this.localColumns.length
      const rowsLen = this.localRows.length
      if (rowsLen && columnsLen) {
        _columns.push(mergeBaseInfo({
          __index: `0-0`,
          colspan: this.localRows.length,
          rowspan: this.localColumns.length,
          isCorner: true
        }))
      }
      this.localRows.forEach(({ label }, index) => {
        _columns.push(mergeBaseInfo({
          __index: `${this.localColumns.length}-${index}`,
          value: label,
          x: this.localColumns.length,
          y: index
        }))
      })
      return _columns
    },
    rowHeadValues() {
      let _values = []
      const rowsLen = this.localRows.length
      if (rowsLen) {
        const rowSpans = {}
        _values = this.rowPaths.map((path, pathIndex) => {
          const values = path.split(this.Separator)
          const currPath = []
          return this.localRows.map((item, rowIndex) => {
            const currVal = values[rowIndex] || ''
            const baseX = this.localColumns.length + +Boolean(this.localValues.length) + pathIndex
            const baseY = rowIndex
            currPath.push(currVal)
            const isLastCol = rowIndex === rowsLen - 1
            if (!isLastCol) {
              const conditions = {}
              for (let i = 0; i < rowIndex + 1; i++) {
                conditions[i] = values[i] || ''
              }
              const filterData = this.rowPaths.filter((data) => {
                let status = true
                const splitValues = data.split(this.Separator)
                for (const key in conditions) {
                  if (conditions[key] !== splitValues[key]) {
                    status = false
                    return
                  }
                }
                return status
              }) || []
              const mergeSpans = filterData.length
              const currRowSpan = rowSpans[rowIndex] || {}
              const currRowSpanVal = (currRowSpan[currPath.join(this.Separator)] || 0) + 1
              currRowSpan[currPath.join(this.Separator)] = currRowSpanVal
              rowSpans[rowIndex] = currRowSpan
              if (currRowSpanVal === 1) {
                return mergeBaseInfo({
                  __index: `${baseX}-${baseY}`,
                  value: currVal,
                  x: baseX,
                  y: baseY,
                  rowspan: mergeSpans,
                  path: currPath.filter((item) => !!item)
                })
              } else {
                return mergeBaseInfo({
                  __index: `${baseX}-${baseY}`,
                  value: currVal,
                  x: baseX,
                  y: baseY,
                  path: currPath.filter((item) => !!item),
                  isRowspan: true
                })
              }
            } else {
              return mergeBaseInfo({
                __index: `${baseX}-${baseY}`,
                value: currVal,
                x: baseX,
                y: baseY,
                path: currPath.filter((item) => !!item)
              })
            }
          })
        })
      }
      return _values
    },
    dataValues() {
      const colConditions = convertPathToMap(
        this.colPaths,
        this.localColumns.map(({ key }) => key).concat(this.localValues.length ? ['value'] : [])
      )
      const rowConditions = convertPathToMap(
        this.rowPaths,
        this.localRows.map(({ key }) => key)
      )
      !colConditions.length && colConditions.push({})
      !rowConditions.length && rowConditions.push({})
      return rowConditions.map((rowCondition, rowConditionIndex) => {
        const _data = colConditions.map((colCondition, colConditionIndex) => {
          const cellData = {}
          const conditions = Object.assign({}, rowCondition, colCondition)
          const _filterConditions = Object.fromEntries(
            Object.entries(conditions).filter(
              (item) => item[0] !== 'value'
            )
          )
          const filterData = this._filterData(_filterConditions, this.localData)
          const baseX = this.localColumns.length + +Boolean(this.localValues.length) + rowConditionIndex
          const baseY = this.localRows.length + colConditionIndex
          Object.assign(
            cellData,
            mergeBaseInfo({
              conditions,
              x: baseX,
              y: baseY,
              __index: `${baseX}-${baseY}`
            })
          )
          const isEmptyValues = this.localColumns.length && this.localRows.length && !this.localValues.length
          if (isEmptyValues) {
            Object.assign(cellData, { value: '' })
          } else {
            const _value = this.values.find(({ key }) => key === conditions.value)
            Object.assign(cellData, { value: _value && _value.key ? this._reduceValue(filterData, _value.key) : '' })
          }
          return cellData
        })
        return {
          __index: _data[0].x,
          data: _data
        }
      })
    },
    trhead() {
      return this.tr.head.filter(function(cell) { return !cell.isRowspan })
    }
  },
  watch: {
    watchAllProps() {
      this.init()
    }
  },
  created() {
    this.init()
  },
  methods: {
    init() {
      if (this.rows.length || this.columns.length || this.values.length) {
        this.handleDataClone()
        this.setValuesToColAndRow()
        this.handleCalcData()
        this.handleCombineHeads()
        this.handleCombineValues()
      } else {
        console.warn('[Warn]: props.rows, props.columns, props.values at least one is not empty.')
      }
    },
    handleDataClone() {
      this.localRows = JSON.parse(JSON.stringify(this.rows))
      this.localColumns = JSON.parse(JSON.stringify(this.columns))
      this.localValues = JSON.parse(JSON.stringify(this.values))
      this.localData = Object.freeze(this.data)
    },
    setValuesToColAndRow() {
      const rowKeys = this.localRows.map(({ key }) => key)
      const columnKeys = this.localColumns.map(({ key }) => key)
      const rowValues = this._findCategory(rowKeys, this.localData)
      const columnValues = this._findCategory(columnKeys, this.localData)
      this.localRows.forEach((row) => {
        const { key, values } = row
        this.$set(row, 'values', values || rowValues[key] || [])
      })
      this.localColumns.forEach((column) => {
        const { key, values } = column
        this.$set(column, 'values', values || columnValues[key] || [])
      })
    },
    handleCombineHeads() {
      let combineColHeads = JSON.parse(JSON.stringify(this.colHeads))
      combineColHeads[0] = combineColHeads[0] || []
      combineColHeads[0].unshift(...this.rowHeads.filter((item) => item.isCorner))
      combineColHeads[combineColHeads.length - 1].unshift(...this.rowHeads.filter((item) => !item.isCorner))
      combineColHeads = combineColHeads.filter((item) => item.length)
      this.combineHeads = combineColHeads
    },
    handleCombineValues() {
      const combineValues = []
      const valueRowCount = this.dataValues.length || this.rowHeadValues.length
      for (let i = 0; i < valueRowCount; i++) {
        const _currRowHeadValue = this.rowHeadValues[i] || []
        const _currValue = this.dataValues[i] || {}
        const _row = [...(_currValue.data || [])]
        combineValues.push(
          Object.assign({}, { head: [..._currRowHeadValue] }, { data: _row })
        )
      }
      this.combineValues = combineValues
    },
    handleCalcData() {
      if (!this.localValues.length) return
      const _rowPaths = this._combineRowPaths(
        this.localData,
        ...this.localRows.map(({ key, values }) => {
          return { key, values }
        })
      )
      const _rowKeys = this.localRows.map(({ key }) => key)
      const _colPaths = this._combineColPaths(
        ...this.localColumns.map(({ values }) => values)
      )
      const _colKeys = this.localColumns.map(({ key }) => key)
      const colConditions = convertPathToMap(_colPaths, _colKeys)
      const rowConditions = convertPathToMap(_rowPaths, _rowKeys)
      !colConditions.length && colConditions.push({})
      !rowConditions.length && rowConditions.push({})
      this.calcData = Object.freeze(
        rowConditions
          .map((rowCondition, rowConditionIndex) =>
            colConditions
              .map((colCondition, colConditionIndex) => {
                const conditions = Object.assign({}, rowCondition, colCondition)
                const filterData = this._filterData(conditions, this.localData)
                const isEmptyCell = this.localRows.length && this.localColumns.length && !this.localValues.length
                const _values = {}
                this.values.forEach(({ key }) => {
                  _values[key] = isEmptyCell ? '' : this._reduceValue(filterData, key)
                })
                return Object.assign({}, conditions, _values)
              })
              .flat()
          )
          .filter((item) => item.length)
          .flat()
      )
    },
    _combineRowPaths(data, ...arrays) {
      const len = arrays.length
      let _result = []
      if (len) {
        const rowPaths = arrays.reduce((prev, curr) => {
          const arr = []
          prev.values.forEach(_prevEl => {
            const prevKey = prev.key.split(SEPARATOR)
            curr.values.forEach(_currEl => {
              const currKey = curr.key
              const conditions = {}
              prevKey.forEach((key, i) => {
                conditions[key] = _prevEl.split(SEPARATOR)[i]
              })
              conditions[currKey] = _currEl
              const filter = data.some((data) => {
                let status = true
                for (const key in conditions) {
                  if (conditions[key] !== data[key]) {
                    status = false
                    return
                  }
                }
                return status
              })
              if (filter) {
                arr.push(_prevEl + SEPARATOR + _currEl)
              }
            })
          })
          return { key: prev.key + SEPARATOR + curr.key, values: arr }
        }) || {}
        _result = rowPaths.values || []
      }
      return _result
    },
    _combineColPaths(...arrays) {
      return arrays.length ? arrays.reduce((prev, curr) => {
        const arr = []
        prev.forEach(_prevEl => {
          curr.forEach(_currEl => {
            arr.push(_prevEl + SEPARATOR + _currEl)
          })
        })
        return arr
      }) : arrays
    },
    _findCategory(keys = [], data = []) {
      const _result = {}
      data.forEach(item => {
        keys.forEach(key => {
          _result[key] = _result[key] || []
          _result[key].push(item[key])
          _result[key] = [...new Set(_result[key])]
        })
      })
      return _result
    },
    _reduceValue(data, key) {
      if (!data.length) return ''
      return data.reduce((sum, item) => {
        return sum + Number(item[key])
      }, 0)
    },
    _filterData(conditions, data) {
      return data.filter((data) => {
        let status = true
        for (const key in conditions) {
          if (conditions[key] !== data[key]) {
            status = false
            return
          }
        }
        return status
      })
    },
    _getMinHeightByRowCount(count) {
      return getHeightByCount(count)
    }
  }
}
</script>

<style lang="scss" scoped>
table {
  border-collapse: collapse;
  border-spacing: 0;
  border: none;

  td, th {
    border: 1px solid #ccc;
    padding: 0;
    vertical-align: middle;
    box-sizing: border-box;

    > div {
      display: flex;
      align-items: center;
      justify-content: center;
      box-sizing: border-box;
      padding: 5px;
      text-align: center;
      white-space: nowrap;
      width: 100%;
      height: 100%;
      min-height: 36px;
      cursor: default;

      &.col-corner-bg {
      }
    }
  }
}
</style>
