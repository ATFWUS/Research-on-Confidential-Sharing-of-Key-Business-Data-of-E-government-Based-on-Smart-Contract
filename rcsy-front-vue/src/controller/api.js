import L from '@/config/axios'
import Axios from 'axios';
import qs from 'qs'

// 登陆
export const login = data => L('/login', data, "POSTJSON");
// 注册
export const register = data => L('/register', data, "POSTJSON");

// 发送信息
export const send = data => L('/send', data, "POSTJSON");

// 搜索信息
export const searchfile = data => L('/receive', data, "POST");

// 搜索信息
export const searchfilee = data => L('/receivee', data, "POST");

export const update = data => L('/send', data, "POSTJSON");
// 退出当前账号
export const logout = data => L('/logout', data, 'GET');

//验证数据完整性
export const handleYan = data =>L('/rehash',data,"POSTJSON")