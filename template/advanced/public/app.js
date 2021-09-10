import {MDCRipple} from '@material/ripple/index';
import {MDCTextField} from '@material/textfield/component';

const ripple = new MDCRipple(document.querySelector('.signin-button'));
const usernameTF = new MDCTextField(document.querySelector('.username-textfield'));
const passwordTF = new MDCTextField(document.querySelector('.password-textfield'));
