import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/requester.dart';

import 'bodies/comment.dart';

Future<HttpRequest> createComment(Comment obj) async {
  var apiroute = getEndpoint("comment");
  var url = "${apiroute}/messages";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}
