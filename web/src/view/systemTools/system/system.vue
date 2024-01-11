<template>
  <div class="system">
    <el-form
        ref="form"
        :inline="true"
        :model="config"
        label-width="120px"
    >
      <el-collapse v-model="activeNames">
        <template v-if="config.system">
          <el-collapse-item name="1" title="系统基本配置">
            <el-form-item label="环境值">
              <el-select v-model="config.system.env" style="width:100%">
                <el-option value="public"/>
                <el-option value="develop"/>
              </el-select>
            </el-form-item>
            <el-form-item label="端口值">
              <el-input v-model.number="config.system.addr"/>
            </el-form-item>
            <el-form-item label="数据库类型">
              <el-select v-model="config.system['db-type']" style="width:100%">
                <el-option value="mysql"/>
                <el-option value="pgsql"/>
              </el-select>
            </el-form-item>
            <el-form-item label="Oss类型">
              <el-select v-model="config.system['oss-type']" style="width:100%">
                <el-option value="local"/>
                <el-option value="qiniu"/>
                <el-option value="tencent-cos"/>
                <el-option value="aliyun-oss"/>
                <el-option value="huawei-obs"/>
              </el-select>
            </el-form-item>
            <el-form-item label="多点登录拦截">
              <el-checkbox v-model="config.system['use-multipoint']">开启</el-checkbox>
            </el-form-item>
            <el-form-item label="开启redis">
              <el-checkbox v-model="config.system['use-redis']">开启</el-checkbox>
            </el-form-item>
            <el-form-item label="限流次数">
              <el-input-number v-model.number="config.system['ip-limit-count']" style="width: 193px;"/>
            </el-form-item>
            <el-form-item label="限流时间">
              <el-input-number v-model.number="config.system['ip-limit-time']" style="width: 193px;"/>
            </el-form-item>
            <el-tooltip
                content="请修改完成后，注意一并修改前端env环境下的VITE_BASE_PATH"
                placement="top-start"
            >
              <el-form-item label="全局路由前缀">
                <el-input v-model="config.system['router-prefix']"/>
              </el-form-item>
            </el-tooltip>
          </el-collapse-item>
        </template>
        <template v-if="config.jwt">
          <el-collapse-item name="2" title="jwt 配置">
            <el-form-item label="jwt签名">
              <el-input v-model="config.jwt['signing-key']" style="width: 300px"/>
            </el-form-item>
            <el-form-item label="有效期">
              <el-input v-model="config.jwt['expires-time']"/>
            </el-form-item>
            <el-form-item label="缓冲期">
              <el-input v-model="config.jwt['buffer-time']"/>
            </el-form-item>
            <el-form-item label="签发者">
              <el-input v-model="config.jwt.issuer"/>
            </el-form-item>
          </el-collapse-item>
        </template>
        <template v-if="config.zap">
          <el-collapse-item name="3" title="Zap 日志配置">
            <el-form-item label="级别">
              <el-input v-model.number="config.zap.level"/>
            </el-form-item>
            <el-form-item label="输出">
              <el-input v-model="config.zap.format"/>
            </el-form-item>
            <el-form-item label="日志前缀">
              <el-input v-model="config.zap.prefix"/>
            </el-form-item>
            <el-form-item label="日志文件夹">
              <el-input v-model="config.zap.director"/>
            </el-form-item>
            <el-form-item label="编码级">
              <el-input v-model="config.zap['encode-level']"/>
            </el-form-item>
            <el-form-item label="栈名">
              <el-input v-model="config.zap['stacktrace-key']"/>
            </el-form-item>
            <el-form-item label="日志留存时间(天)">
              <el-input v-model.number="config.zap['max-age']"/>
            </el-form-item>
            <el-form-item label="显示行">
              <el-checkbox v-model="config.zap['show-line']"/>
            </el-form-item>
            <el-form-item label="输出控制台">
              <el-checkbox v-model="config.zap['log-in-console']"/>
            </el-form-item>
          </el-collapse-item>
        </template>
        <template v-if="config.local">
          <el-collapse-item name="4" title="文件上传配置">
            <template v-if="config.system['oss-type'] === 'local'">
              <h2>本地文件配置</h2>
              <el-form-item label="文件访问路径">
                <el-input v-model="config.local.path"/>
              </el-form-item>
              <el-form-item label="文件存储路径">
                <el-input v-model="config.local['store-path']"/>
              </el-form-item>
            </template>
            <template v-if="config.system['oss-type'] === 'qiniu'">
              <h2>qiniu上传配置</h2>
              <el-form-item label="存储区域">
                <el-input v-model="config.qiniu.zone"/>
              </el-form-item>
              <el-form-item label="空间名称">
                <el-input v-model="config.qiniu.bucket"/>
              </el-form-item>
              <el-form-item label="CDN加速域名">
                <el-input v-model="config.qiniu['img-path']"/>
              </el-form-item>
              <el-form-item label="是否使用https">
                <el-checkbox v-model="config.qiniu['use-https']">开启</el-checkbox>
              </el-form-item>
              <el-form-item label="accessKey">
                <el-input v-model="config.qiniu['access-key']"/>
              </el-form-item>
              <el-form-item label="secretKey">
                <el-input v-model="config.qiniu['secret-key']"/>
              </el-form-item>
              <el-form-item label="上传是否使用CDN上传加速">
                <el-checkbox v-model="config.qiniu['use-cdn-domains']">开启</el-checkbox>
              </el-form-item>
            </template>
            <template v-if="config.system['oss-type'] === 'tencent-cos'">
              <h2>腾讯云COS上传配置</h2>
              <el-form-item label="存储桶名称">
                <el-input v-model="config['tencent-cos']['bucket']"/>
              </el-form-item>
              <el-form-item label="所属地域">
                <el-input v-model="config['tencent-cos'].region"/>
              </el-form-item>
              <el-form-item label="secretID">
                <el-input v-model="config['tencent-cos']['secret-id']"/>
              </el-form-item>
              <el-form-item label="secretKey">
                <el-input v-model="config['tencent-cos']['secret-key']"/>
              </el-form-item>
              <el-form-item label="路径前缀">
                <el-input v-model="config['tencent-cos']['path-prefix']"/>
              </el-form-item>
              <el-form-item label="访问域名">
                <el-input v-model="config['tencent-cos']['base-url']"/>
              </el-form-item>
            </template>

          </el-collapse-item>
        </template>
        <template v-if="config.excel">
          <el-collapse-item name="5" title="Excel 上传配置">
            <el-form-item label="合成目标地址">
              <el-input v-model="config.excel.dir"/>
            </el-form-item>
          </el-collapse-item>
        </template>
        <template v-if="config.timer">
          <el-collapse-item name="6" title="定时任务配置">
            <el-form-item label="CRON表达式">
              <el-input v-model="config.timer.spec"/>
            </el-form-item>
            <template v-for="(item,k) in config.timer.detail">
              <div v-for="(_,k2) in item" :key="k2">
                <el-form-item :key="k+k2" :label="k2">
                  <el-input v-model="item[k2]"/>
                </el-form-item>
              </div>
            </template>
            <el-form-item label="是否启用">
              <el-checkbox v-model="config.timer['start']"/>
            </el-form-item>
          </el-collapse-item>
        </template>
        <template v-if="config.redis">
          <el-collapse-item name="7" title="Redis 数据库配置">
            <el-form-item label="库">
              <el-input v-model.number="config.redis.db"/>
            </el-form-item>
            <el-form-item label="地址">
              <el-input v-model="config.redis.addr"/>
            </el-form-item>
            <el-form-item label="密码">
              <el-input v-model="config.redis.password"/>
            </el-form-item>
          </el-collapse-item>
        </template>
        <template v-if="config.gormDB">
          <el-collapse-item name="8" title="系统数据库配置">
            <template v-if="config.gormDB['db-type'] === 'mysql'">
              <el-form-item label="数据库类型">
                <el-input v-model="config.gormDB['db-type']"/>
              </el-form-item>
              <el-form-item label="ip地址">
                <el-input v-model="config.gormDB.ip"/>
              </el-form-item>
              <el-form-item label="ip端口">
                <el-input v-model="config.gormDB.port"/>
              </el-form-item>
              <el-form-item label="用户名">
                <el-input v-model="config.gormDB.username"/>
              </el-form-item>
              <el-form-item label="密码">
                <el-input v-model="config.gormDB.password"/>
              </el-form-item>
              <el-form-item label="数据库名">
                <el-input v-model="config.gormDB['db-name']"/>
              </el-form-item>
              <el-form-item label="maxIdleConns">
                <el-input v-model.number="config.gormDB['max-idle-conns']"/>
              </el-form-item>
              <el-form-item label="maxOpenConns">
                <el-input v-model.number="config.gormDB['max-open-conns']"/>
              </el-form-item>
              <el-form-item label="日志模式">
                <el-input v-model="config.gormDB['log-mode']"/>
              </el-form-item>
              <el-form-item label="写入日志">
                <el-checkbox v-model="config.gormDB['log-zap']"/>
              </el-form-item>
            </template>
          </el-collapse-item>
        </template>
        <template v-if="config.ods_db">
          <el-collapse-item name="9" title="ODS数据库配置">
            <template v-if="config.system['db-type'] === 'mysql'">
              <el-form-item label="数据库类型">
                <el-input v-model="config.ods_db['db-type']"/>
              </el-form-item>
              <el-form-item label="ip地址">
                <el-input v-model="config.ods_db.ip"/>
              </el-form-item>
              <el-form-item label="ip端口">
                <el-input v-model="config.ods_db.port"/>
              </el-form-item>
              <el-form-item label="用户名">
                <el-input v-model="config.ods_db.username"/>
              </el-form-item>
              <el-form-item label="密码">
                <el-input v-model="config.ods_db.password"/>
              </el-form-item>
              <el-form-item label="数据库名">
                <el-input v-model="config.ods_db['db-name']"/>
              </el-form-item>
              <el-form-item label="maxIdleConns">
                <el-input v-model.number="config.ods_db['max-idle-conns']"/>
              </el-form-item>
              <el-form-item label="maxOpenConns">
                <el-input v-model.number="config.ods_db['max-open-conns']"/>
              </el-form-item>
              <el-form-item label="日志模式">
                <el-input v-model="config.ods_db['log-mode']"/>
              </el-form-item>
              <el-form-item label="写入日志">
                <el-checkbox v-model="config.ods_db['log-zap']"/>
              </el-form-item>
            </template>
          </el-collapse-item>
        </template>
        <template v-if="config.uds_db">
          <el-collapse-item name="10" title="UDS数据库配置">
            <template v-if="config.system['db-type'] === 'mysql'">
              <el-form-item label="数据库类型">
                <el-input v-model="config.uds_db['db-type']"/>
              </el-form-item>
              <el-form-item label="ip地址">
                <el-input v-model="config.uds_db.ip"/>
              </el-form-item>
              <el-form-item label="ip端口">
                <el-input v-model="config.uds_db.port"/>
              </el-form-item>
              <el-form-item label="用户名">
                <el-input v-model="config.uds_db.username"/>
              </el-form-item>
              <el-form-item label="密码">
                <el-input v-model="config.uds_db.password"/>
              </el-form-item>
              <el-form-item label="数据库名">
                <el-input v-model="config.uds_db['db-name']"/>
              </el-form-item>
              <el-form-item label="maxIdleConns">
                <el-input v-model.number="config.uds_db['max-idle-conns']"/>
              </el-form-item>
              <el-form-item label="maxOpenConns">
                <el-input v-model.number="config.uds_db['max-open-conns']"/>
              </el-form-item>
              <el-form-item label="日志模式">
                <el-input v-model="config.uds_db['log-mode']"/>
              </el-form-item>
              <el-form-item label="写入日志">
                <el-checkbox v-model="config.uds_db['log-zap']"/>
              </el-form-item>
            </template>
          </el-collapse-item>
        </template>
        <template v-if="config.etcd">
          <el-collapse-item name="11" title="系统 ETCD 配置">
            <el-form-item label="dsn">
              <el-input v-model="config.etcd.dsn"/>
            </el-form-item>
          </el-collapse-item>
        </template>
        <template v-if="config.kafka">
          <el-collapse-item name="12" title="系统 kafka 配置">
            <el-form-item label="dsn">
              <el-input v-model="config.kafka.dsn"/>
            </el-form-item>
          </el-collapse-item>
        </template>
        <template v-if="config.canal">
          <el-collapse-item name="13" title="系统 canal 配置">
            <el-form-item label="mysql-dump" style="width: 800px;">
              <el-input v-model="config.canal.mysqlDump"/>
            </el-form-item>
          </el-collapse-item>
        </template>
      </el-collapse>
    </el-form>
    <div class="gva-btn-list">
      <el-button type="primary" @click="update">立即更新</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Config'
}
</script>
<script setup>
import {getSystemConfig, setSystemConfig} from '@/api/system'
import {ref, reactive} from 'vue'
import {ElMessage} from 'element-plus'

const activeNames = reactive([])
const config = ref({
  system: {
    'ip-limit-count': 15000,
    'ip-limit-time': 3600
  },
  jwt: {},
  mysql: {},
  pgsql: {},
  excel: {},
  redis: {},
  captcha: {},
  zap: {},
  local: {},
  email: {},
  timer: {
    detail: {}
  },
  sys_db: {},
  ods_db: {},
  uds_db: {},
  etcd: {},
  kafka: {},
  canal: {},
})

const initForm = async () => {
  const res = await getSystemConfig()
  if (res.code === 0) {
    config.value = res.data.config
  }
}
initForm()
const reload = () => {
}
const update = async () => {
  const res = await setSystemConfig({config: config.value})
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '配置文件设置成功'
    })
    await initForm()
  }
}
</script>

<style lang="scss">
.system {
  background: #fff;
  padding: 36px;
  border-radius: 2px;

  h2 {
    padding: 10px;
    margin: 10px 0;
    font-size: 16px;
    box-shadow: -4px 0px 0px 0px #e7e8e8;
  }

  ::v-deep(.el-input-number__increase) {
    top: 5px !important;
  }

  .gva-btn-list {
    margin-top: 16px;
  }
}

.el-form-item {
  margin-right: 0 !important;
}

.el-form-item__label {
  position: absolute;
}

.el-form-item__content {
  width: 100%;
  padding-left: 120px;
}

.el-select {
  width: 193px !important;
}

.el-input_inner {
  width: 100%;
}
</style>
