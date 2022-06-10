/**
 * 存储localStorage
 */
export const setStore = (name, content) => {
  if (!name) return;
  if (typeof content !== 'string') {
    content = JSON.stringify(content);
  }
  window.localStorage.setItem(name, content);
};

/**
 * 获取localStorage
 */
export const getStore = name => {
  if (!name) return;
  return window.localStorage.getItem(name);
};

/**
 * 删除localStorage
 */
export const removeStore = name => {
  if (!name) return;
  window.localStorage.removeItem(name);
};

/**
 * 日期格式化
 * @param date
 * @param fmt
 * @returns {*}
 */
export function formatDate(date, fmt) {
  if(typeof(date) === "number"){
    date = new Date(date);
  }
  if (/(y+)/.test(fmt)) {
    fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length));
  }
  let o = {
    'M+': date.getMonth() + 1,
    'd+': date.getDate(),
    'h+': date.getHours(),
    'm+': date.getMinutes(),
    's+': date.getSeconds()
  };
  for (let k in o) {
    if (new RegExp(`(${k})`).test(fmt)) {
      let str = o[k] + '';
      fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? str : padLeftZero(str));
    }
  }
  return fmt;
};

/**
 * 状态
 */
export function taskStatus(status){
  if(status === 0){
    return '未收到'
  }else if(status === 1){
    return '进行中'
  }else if(status === 2){
    return '已完毕'
  }else if(status === 3){
    return '已完成'
  }
}

export function fileType(file){
  let index = file.lastIndexOf('.')
  let str = file.substring(index+1,file.length)
  return str
}

export function getBasePath(){
  return process.env.API_PATH?process.env.API_PATH:"";
}
function padLeftZero (str) {
  return ('00' + str).substr(str.length);
};
