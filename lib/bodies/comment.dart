import 'package:mango_ui/keys.dart';

class Comment {
  final Key itemKey;
  final String text;
  final num commentType;
  final String userImage;

  List<Comment> _children;

  Comment(this.itemKey, this.text, this.commentType, this.userImage);

  void addChild(Comment child) {
    _children.add(child);
  }

  Map<String, dynamic> toJson() {
    return {
      "ItemKey": itemKey,
      "Text": text,
      "CommentType": commentType,
      "Children": _children,
      "UserImage": userImage,
    };
  }
}
