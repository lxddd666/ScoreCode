import {useGlobSetting} from "@/hooks/setting";
import {storage} from "@/utils/Storage";
import {ACCESS_TOKEN} from "@/store/mutation-types";

const globSetting = useGlobSetting();
const tgPrefix = globSetting.tgPrefix || '';

const token = storage.get(ACCESS_TOKEN);


export function getPhoto(account: any, user: any, photoId: number): string {
  return photoId == 0 ? 'https://gw.alipayobjects.com/zos/antfincdn/aPkFc8Sj7n/method-draw-image.svg' :
    tgPrefix + '/arts/user/getUserAvatar?authorization=' + token + '&account=' + account + '&getUser=' + user + '&photoId=' + photoId
}
